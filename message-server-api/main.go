package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-martini/martini"
	"github.com/go-playground/validator"
	messagecommonlib "github.com/hectorandac/grpc-poc/message-common-lib"
	"github.com/hectorandac/grpc-poc/message-server-api/middlewares"
	pb "github.com/hectorandac/grpc-poc/protos"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var validate *validator.Validate
var protoClient pb.BillingClient

const (
	address = "localhost:50051"
	secrete = "64F5AADE234AA8E534DBDE5DCB5C5"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}
	defer conn.Close()
	protoClient = pb.NewBillingClient(conn)

	validate = validator.New()
	m := martini.Classic()

	m.Use(middlewares.MongoDB())
	m.Use(render.Renderer())

	m.Post("/message", binding.Bind(messagecommonlib.Message{}), processMessage)
	m.Post("/refund", binding.Bind(messagecommonlib.RefundEntity{}), processRefund)
	m.Post("/message-n", binding.Bind(messagecommonlib.Message{}), processMessageN)
	m.Post("/refund-n", binding.Bind(messagecommonlib.RefundEntity{}), processRefund)

	m.RunOnAddr(":3000")
}

func processMessage(message messagecommonlib.Message, r render.Render, db *mgo.Database, req *http.Request) {
	message.ReceivedOn = time.Now().UnixNano()
	message.Refunded = false
	validationError := validate.Struct(message)

	if validationError != nil {
		json := map[string]interface{}{"error": strings.Split(validationError.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	token := req.Header.Get("Authorization")
	claims, tokenErr := extractClaims(token)
	if token == "" || tokenErr != nil {
		json := map[string]interface{}{"error": strings.Split(tokenErr.Error(), "\n")}
		r.JSON(403, json)
		return
	}

	var balance float32 = 0
	if claims["balance"] != nil {
		balance = float32(claims["balance"].(float64))
	}

	billableEntity := &pb.BillableEntity{
		FullName:         claims["fullname"].(string),
		BillableEntityID: claims["billableentityid"].(string),
		UserUUID:         claims["useruuid"].(string),
		Balance:          balance,
	}

	commonCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	billableEntityId, _ := protoClient.FindOrCreateBillingEntity(commonCtx, billableEntity)
	if !billableEntityId.Success {
		json := map[string]interface{}{"error": "Unable to create or retrieve billing infor for user"}
		r.JSON(400, json)
		return
	}

	message.Id = bson.ObjectId(randomHex(6))
	billingRequest := &pb.BillRequest{
		UserUUID: claims["useruuid"].(string),
		Deliverable: &pb.Message{
			Id:          message.Id.Hex(),
			Recipient:   message.Recipient,
			Body:        message.Message,
			Sender:      message.Sender,
			Type:        message.Type,
			CreatedOn:   message.CreatedOn,
			ReceivedOn:  message.ReceivedOn,
			ProcessedOn: message.ProcessedOn,
			Refunded:    message.Refunded,
		},
	}

	messagePrice, _ := protoClient.CalculateMessageRate(commonCtx, billingRequest)
	if !messagePrice.Success {
		json := map[string]interface{}{"error": "Unable to retrieve messaging price"}
		r.JSON(400, json)
		return
	}

	userBalance, _ := protoClient.UserBalance(commonCtx, billableEntity)
	if !userBalance.Success {
		json := map[string]interface{}{"error": "Unable to retrieve user balance"}
		r.JSON(400, json)
		return
	}

	userBalanceNumerical, _ := strconv.ParseFloat(userBalance.GetBody(), 32)
	messagePriceNumerical, _ := strconv.ParseFloat(messagePrice.GetBody(), 32)
	if (userBalanceNumerical - messagePriceNumerical) < 0 {
		json := map[string]interface{}{"error": "You don't have enough balance for this message"}
		r.JSON(400, json)
		return
	}

	message.Rate = messagePriceNumerical
	billingResponse, _ := protoClient.BillUser(commonCtx, billingRequest)
	if !billingResponse.Success {
		json := map[string]interface{}{"error": "Couldn't bill you for this message", "body": billingResponse.GetBody()}
		r.JSON(400, json)
		return
	}

	message.CreatedOn = time.Now().UnixNano()
	message.ProcessTime = message.CreatedOn - message.ReceivedOn
	err := db.C("messages").Insert(message)
	if err != nil {
		json := map[string]interface{}{"error": strings.Split(err.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	message.ProcessedOn = time.Now().UnixNano()
	updateErr := db.C("messages").UpdateId(message.Id, message)
	if updateErr != nil {
		json := map[string]interface{}{"error": strings.Split(updateErr.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	json := map[string]interface{}{
		"status":          "success",
		"message":         message,
		"message_charge":  messagePrice.GetBody(),
		"current_balance": userBalanceNumerical - messagePriceNumerical,
		"origin":          claims,
	}
	r.JSON(200, json)
}

func extractClaims(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secrete), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func processRefund(message_id messagecommonlib.RefundEntity, r render.Render, db *mgo.Database, req *http.Request) {

	token := req.Header.Get("Authorization")
	claims, tokenErr := extractClaims(token)
	if token == "" || tokenErr != nil {
		json := map[string]interface{}{"error": strings.Split(tokenErr.Error(), "\n")}
		r.JSON(403, json)
		return
	}

	var balance float32 = 0
	if claims["balance"] != nil {
		balance = float32(claims["balance"].(float64))
	}

	billableEntity := &pb.BillableEntity{
		FullName:         claims["fullname"].(string),
		BillableEntityID: claims["billableentityid"].(string),
		UserUUID:         claims["useruuid"].(string),
		Balance:          balance,
	}

	commonCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	billableEntityId, _ := protoClient.FindOrCreateBillingEntity(commonCtx, billableEntity)
	if !billableEntityId.Success {
		json := map[string]interface{}{"error": "Unable to create or retrieve billing info for user"}
		r.JSON(400, json)
		return
	}

	userBalance, _ := protoClient.UserBalance(commonCtx, billableEntity)
	if !userBalance.Success {
		json := map[string]interface{}{"error": "Unable to retrieve user balance"}
		r.JSON(400, json)
		return
	}

	t, _ := strconv.ParseFloat(userBalance.Body, 32)
	billableEntity.Balance = float32(t)

	result := messagecommonlib.Message{}
	err := db.C("messages").Find(bson.M{"_id": bson.ObjectIdHex(message_id.Id)}).One(&result)
	if err != nil {
		json := map[string]interface{}{"error": strings.Split(err.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	message_refund := &pb.Message{
		Id:          result.Id.Hex(),
		Recipient:   result.Recipient,
		Body:        result.Message,
		Sender:      result.Sender,
		Type:        result.Type,
		CreatedOn:   result.CreatedOn,
		ReceivedOn:  result.ReceivedOn,
		ProcessedOn: result.ProcessedOn,
		Rate:        float32(result.Rate),
		Refunded:    result.Refunded,
	}

	refund_request := &pb.RefundRequest{
		Message:        message_refund,
		BillableEntity: billableEntity,
	}

	updatedBillableEntity, err := protoClient.RefundUser(commonCtx, refund_request)
	if !updatedBillableEntity.Success {
		json := map[string]interface{}{"error": updatedBillableEntity.Body}
		r.JSON(400, json)
		return
	}

	update_message := bson.M{
		"$set": messagecommonlib.Message{
			Id:          bson.ObjectIdHex(refund_request.GetMessage().GetId()),
			Recipient:   refund_request.GetMessage().Recipient,
			Message:     refund_request.GetMessage().GetBody(),
			Sender:      refund_request.GetMessage().Sender,
			Type:        refund_request.GetMessage().Type,
			CreatedOn:   refund_request.GetMessage().CreatedOn,
			ReceivedOn:  refund_request.GetMessage().ReceivedOn,
			ProcessedOn: refund_request.GetMessage().ProcessedOn,
			Rate:        float64(refund_request.GetMessage().Rate),
			Refunded:    true,
		},
	}
	db.C("messages").Update(bson.M{"_id": bson.ObjectIdHex(refund_request.GetMessage().Id)}, update_message)

	json := map[string]interface{}{
		"status":          "success",
		"message":         message_id.Id + " refunded successfully",
		"balance_details": updatedBillableEntity.Body,
		"origin":          claims,
	}
	r.JSON(200, json)
}

func processMessageN(message messagecommonlib.Message, r render.Render, db *mgo.Database, req *http.Request) {
	message.ReceivedOn = time.Now().UnixNano()
	message.Refunded = false
	validationError := validate.Struct(message)

	if validationError != nil {
		json := map[string]interface{}{"error": strings.Split(validationError.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	token := req.Header.Get("Authorization")
	claims, tokenErr := extractClaims(token)
	if token == "" || tokenErr != nil {
		json := map[string]interface{}{"error": strings.Split(tokenErr.Error(), "\n")}
		r.JSON(403, json)
		return
	}

	var balance float32 = 0
	if claims["balance"] != nil {
		balance = float32(claims["balance"].(float64))
	}

	billableEntity := &pb.BillableEntity{
		FullName:         claims["fullname"].(string),
		BillableEntityID: claims["billableentityid"].(string),
		UserUUID:         claims["useruuid"].(string),
		Balance:          balance,
	}

	bJson, _ := json.Marshal(billableEntity)
	billableEntityId := httpRequest(
		"GET",
		"http://0.0.0.0:3010/billlingentity",
		strings.NewReader(string(bJson)),
	)
	var bResponse pb.BillResponse
	json.Unmarshal([]byte(billableEntityId), &bResponse)

	if !bResponse.Success {
		json := map[string]interface{}{"error": "Unable to create or retrieve billing infor for user"}
		r.JSON(400, json)
		return
	}

	message.Id = bson.ObjectId(randomHex(6))
	billingRequest := &pb.BillRequest{
		UserUUID: claims["useruuid"].(string),
		Deliverable: &pb.Message{
			Id:          message.Id.Hex(),
			Recipient:   message.Recipient,
			Body:        message.Message,
			Sender:      message.Sender,
			Type:        message.Type,
			CreatedOn:   message.CreatedOn,
			ReceivedOn:  message.ReceivedOn,
			ProcessedOn: message.ProcessedOn,
			Refunded:    message.Refunded,
		},
	}

	cJson, _ := json.Marshal(billingRequest)
	messagePrice := httpRequest(
		"GET",
		"http://0.0.0.0:3010/message-rate",
		strings.NewReader(string(cJson)),
	)
	var cResponse pb.BillResponse
	json.Unmarshal([]byte(messagePrice), &cResponse)
	if !cResponse.Success {
		json := map[string]interface{}{"error": "Unable to retrieve messaging price"}
		r.JSON(400, json)
		return
	}

	dJson, _ := json.Marshal(billableEntity)
	userBalance := httpRequest(
		"GET",
		"http://0.0.0.0:3010/balance",
		strings.NewReader(string(dJson)),
	)
	var dResponse pb.BillResponse
	json.Unmarshal([]byte(userBalance), &dResponse)
	if !dResponse.Success {
		json := map[string]interface{}{"error": "Unable to retrieve user balance"}
		r.JSON(400, json)
		return
	}

	userBalanceNumerical, _ := strconv.ParseFloat(dResponse.GetBody(), 32)
	messagePriceNumerical, _ := strconv.ParseFloat(cResponse.GetBody(), 32)
	if (userBalanceNumerical - messagePriceNumerical) < 0 {
		json := map[string]interface{}{"error": "You don't have enough balance for this message"}
		r.JSON(400, json)
		return
	}

	message.Rate = messagePriceNumerical

	eJson, _ := json.Marshal(billingRequest)
	billingResponse := httpRequest(
		"POST",
		"http://0.0.0.0:3010/bill",
		strings.NewReader(string(eJson)),
	)
	var eResponse pb.BillResponse
	json.Unmarshal([]byte(billingResponse), &eResponse)
	if !eResponse.Success {
		json := map[string]interface{}{"error": "Couldn't bill you for this message", "body": eResponse.GetBody()}
		r.JSON(400, json)
		return
	}

	message.CreatedOn = time.Now().UnixNano()
	message.ProcessTime = message.CreatedOn - message.ReceivedOn
	err := db.C("messages").Insert(message)
	if err != nil {
		json := map[string]interface{}{"error": strings.Split(err.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	message.ProcessedOn = time.Now().UnixNano()
	updateErr := db.C("messages").UpdateId(message.Id, message)
	if updateErr != nil {
		json := map[string]interface{}{"error": strings.Split(updateErr.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	json := map[string]interface{}{
		"status":          "success",
		"message":         message,
		"message_charge":  cResponse.GetBody(),
		"current_balance": userBalanceNumerical - messagePriceNumerical,
		"origin":          claims,
	}
	r.JSON(200, json)
}

func httpRequest(method string, url string, payload io.Reader) string {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

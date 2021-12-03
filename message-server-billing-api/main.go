package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	messagecommonlib "github.com/hectorandac/grpc-poc/message-common-lib"
	"github.com/hectorandac/grpc-poc/message-server-billing/utility"
	pb "github.com/hectorandac/grpc-poc/protos"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var dbClient *mongo.Client
var db *mongo.Database

const (
	messageRate    = 1.21
	messageMaxSize = 24
)

func findOrCreateBillingEntity(billing_entitie pb.BillableEntity, r render.Render, req *http.Request) {
	userInfo := messagecommonlib.BillingEntity{
		UserUUID: billing_entitie.GetUserUUID(),
		Balance:  billing_entitie.GetBalance(),
		FullName: billing_entitie.GetFullName(),
	}
	collection := db.Collection("billing_entities")

	result := messagecommonlib.BillingEntity{}
	err := collection.FindOne(context.TODO(), bson.M{"user_uuid": billing_entitie.GetUserUUID()}).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json := map[string]interface{}{"success": true, "body": result.Id.Hex()}
		r.JSON(200, json)
		return
	}

	newUser, err := collection.InsertOne(context.TODO(), userInfo)
	if err != nil {
		json := map[string]interface{}{"success": false, "body": err.Error()}
		r.JSON(400, json)
		return
	} else {
		json := map[string]interface{}{"success": true, "body": fmt.Sprintf("%x", newUser.InsertedID)}
		r.JSON(200, json)
		return
	}
}

func calculateMessageRate(billing_request pb.BillRequest, r render.Render, req *http.Request) {
	messagePrice := calculateRate(billing_request)
	json := map[string]interface{}{"success": true, "body": fmt.Sprintf("%f", messagePrice)}
	r.JSON(200, json)
	return
}

func userBalance(billing_entitie pb.BillableEntity, r render.Render, req *http.Request) {
	collection := db.Collection("billing_entities")

	result := messagecommonlib.BillingEntity{}
	err := collection.FindOne(context.TODO(), bson.M{"user_uuid": billing_entitie.GetUserUUID()}).Decode(&result)
	if err != nil {
		json := map[string]interface{}{"success": false, "body": err.Error()}
		r.JSON(400, json)
		return
	} else {
		json := map[string]interface{}{"success": true, "body": fmt.Sprintf("%f", result.Balance)}
		r.JSON(200, json)
		return
	}
}

func billUser(billing_request pb.BillRequest, r render.Render, req *http.Request) {
	collection := db.Collection("billing_entities")

	result := messagecommonlib.BillingEntity{}
	err := collection.FindOne(context.TODO(), bson.M{"user_uuid": billing_request.GetUserUUID()}).Decode(&result)
	if err != nil {
		json := map[string]interface{}{"success": false, "body": err.Error()}
		r.JSON(400, json)
		return
	} else {
		newBallance := result.Balance - float32(calculateRate(billing_request))
		result.Balance = newBallance
		update := bson.M{
			"$set": messagecommonlib.BillingEntity{
				UserUUID: result.UserUUID,
				Balance:  newBallance,
				FullName: result.FullName,
			},
		}
		collection.UpdateOne(context.TODO(), bson.M{"user_uuid": billing_request.GetUserUUID()}, update)

		json := map[string]interface{}{"success": true, "body": fmt.Sprintf("%f", result.Balance)}
		r.JSON(200, json)
		return
	}
}

func refundUser(refund_request pb.RefundRequest, r render.Render, req *http.Request) {
	c_billing := db.Collection("billing_entities")

	if refund_request.GetMessage().GetRefunded() {
		json := map[string]interface{}{"success": false, "body": "This message has already been refunded."}
		r.JSON(400, json)
		return
	}

	updated_balance := refund_request.GetBillableEntity().GetBalance() + refund_request.GetMessage().GetRate()
	update_billing := bson.M{
		"$set": messagecommonlib.BillingEntity{
			UserUUID: refund_request.GetBillableEntity().GetUserUUID(),
			Balance:  updated_balance,
			FullName: refund_request.GetBillableEntity().GetFullName(),
		},
	}
	c_billing.UpdateOne(context.TODO(), bson.M{"user_uuid": refund_request.GetBillableEntity().GetUserUUID()}, update_billing)

	json := map[string]interface{}{"success": true, "body": fmt.Sprintf("%f", updated_balance)}
	r.JSON(200, json)
	return
}

func main() {
	dbClient = utility.MongoDB()
	db = dbClient.Database("messsage-server-billing")
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/billlingentity", binding.Bind(pb.BillableEntity{}), findOrCreateBillingEntity)
	m.Get("/message-rate", binding.Bind(pb.BillRequest{}), calculateMessageRate)
	m.Get("/balance", binding.Bind(pb.BillableEntity{}), userBalance)
	m.Post("/bill", binding.Bind(pb.BillRequest{}), billUser)
	m.Post("/refurnd", binding.Bind(pb.RefundRequest{}), refundUser)

	m.RunOnAddr(":3010")
}

func calculateRate(in pb.BillRequest) float64 {
	messageLength := len(in.GetDeliverable().GetBody())
	messageSections := int(messageLength / messageMaxSize)
	return float64(messageSections) * messageRate
}

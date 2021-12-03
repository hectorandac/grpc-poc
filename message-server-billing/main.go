package main

import (
	"context"
	"fmt"
	"log"
	"net"

	messagecommonlib "github.com/hectorandac/grpc-poc/message-common-lib"
	"github.com/hectorandac/grpc-poc/message-server-billing/utility"
	pb "github.com/hectorandac/grpc-poc/protos"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

var dbClient *mongo.Client
var db *mongo.Database

const (
	port           = ":50051"
	messageRate    = 1.21
	messageMaxSize = 24
)

type server struct {
	pb.UnimplementedBillingServer
}

func (s *server) FindOrCreateBillingEntity(ctx context.Context, in *pb.BillableEntity) (*pb.BillResponse, error) {
	userInfo := messagecommonlib.BillingEntity{
		UserUUID: in.GetUserUUID(),
		Balance:  in.GetBalance(),
		FullName: in.GetFullName(),
	}
	collection := db.Collection("billing_entities")

	result := messagecommonlib.BillingEntity{}
	err := collection.FindOne(ctx, bson.M{"user_uuid": in.GetUserUUID()}).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		return &pb.BillResponse{
			Success: true,
			Body:    result.Id.Hex(),
		}, nil
	}

	newUser, err := collection.InsertOne(ctx, userInfo)
	if err != nil {
		return &pb.BillResponse{
			Success: false,
			Body:    err.Error(),
		}, nil
	} else {
		return &pb.BillResponse{
			Success: true,
			Body:    fmt.Sprintf("%x", newUser.InsertedID),
		}, nil
	}
}

func (s *server) CalculateMessageRate(ctx context.Context, in *pb.BillRequest) (*pb.BillResponse, error) {
	messagePrice := calculateRate(in)
	return &pb.BillResponse{
		Success: true,
		Body:    fmt.Sprintf("%f", messagePrice),
	}, nil
}

func (s *server) UserBalance(ctx context.Context, in *pb.BillableEntity) (*pb.BillResponse, error) {
	collection := db.Collection("billing_entities")

	result := messagecommonlib.BillingEntity{}
	err := collection.FindOne(ctx, bson.M{"user_uuid": in.GetUserUUID()}).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return &pb.BillResponse{
			Success: false,
			Body:    err.Error(),
		}, nil
	} else {
		return &pb.BillResponse{
			Success: true,
			Body:    fmt.Sprintf("%f", result.Balance),
		}, nil
	}
}

func (s *server) BillUser(ctx context.Context, in *pb.BillRequest) (*pb.BillResponse, error) {
	collection := db.Collection("billing_entities")

	result := messagecommonlib.BillingEntity{}
	err := collection.FindOne(ctx, bson.M{"user_uuid": in.GetUserUUID()}).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return &pb.BillResponse{
			Success: false,
			Body:    err.Error(),
		}, nil
	} else {
		newBallance := result.Balance - float32(calculateRate(in))
		result.Balance = newBallance
		update := bson.M{
			"$set": messagecommonlib.BillingEntity{
				UserUUID: result.UserUUID,
				Balance:  newBallance,
				FullName: result.FullName,
			},
		}
		collection.UpdateOne(ctx, bson.M{"user_uuid": in.GetUserUUID()}, update)

		return &pb.BillResponse{
			Success: true,
			Body:    fmt.Sprintf("%f", result.Balance),
		}, nil
	}
}

func (s *server) RefundUser(ctx context.Context, in *pb.RefundRequest) (*pb.BillResponse, error) {
	c_billing := db.Collection("billing_entities")

	if in.GetMessage().GetRefunded() {
		return &pb.BillResponse{
			Success: false,
			Body:    "This message has already been refunded.",
		}, nil
	}

	updated_balance := in.GetBillableEntity().GetBalance() + in.GetMessage().GetRate()
	update_billing := bson.M{
		"$set": messagecommonlib.BillingEntity{
			UserUUID: in.GetBillableEntity().GetUserUUID(),
			Balance:  updated_balance,
			FullName: in.GetBillableEntity().GetFullName(),
		},
	}
	c_billing.UpdateOne(ctx, bson.M{"user_uuid": in.GetBillableEntity().GetUserUUID()}, update_billing)

	return s.UserBalance(ctx, in.GetBillableEntity())
}

func main() {
	dbClient = utility.MongoDB()
	db = dbClient.Database("messsage-server-billing")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBillingServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func calculateRate(in *pb.BillRequest) float64 {
	messageLength := len(in.GetDeliverable().GetBody())
	messageSections := int(messageLength / messageMaxSize)
	return float64(messageSections) * messageRate
}

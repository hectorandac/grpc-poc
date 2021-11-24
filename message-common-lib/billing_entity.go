package messagecommonlib

import "gopkg.in/mgo.v2/bson"

type BillingEntity struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	UserUUID string        `json:"user_uuid" form:"user_uuid" binding:"required" bson:"user_uuid"`
	Balance  float32       `json:"balance" bson:"balance"`
	FullName string        `json:"full_name" form:"full_name" binding:"required" bson:"full_name"`
}

package messagecommonlib

type RefundEntity struct {
	Id string `json:"message_id" form:"message_id" binding:"required" bson:"message_id"`
}

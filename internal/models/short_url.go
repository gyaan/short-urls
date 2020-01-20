package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ShortUrl struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Url           string             `json:"actual_url" bson:"url"`
	NewUrl        string             `json:"new_url" bson:"new_url"`
	UrlIdentifier int64              `json:"url_identifier" bson:"url_identifier"`
	ClicksCount   int32              `json:"clicks_count" bson:"clicks_count"`
	ExpireTime    time.Time          `json:"expire_time" bson:"expire_time"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	Status        int32              `json:"status" bson:"status"` //0 - active, 1 - inactive
}

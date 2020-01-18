package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Counter struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Sequence int64 `json:"sequence" bson:"sequence"`
}

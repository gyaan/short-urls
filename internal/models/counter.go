package models

type Counter struct {
	Name     string `json:"name" bson:"name"`
	Sequence int64 `json:"sequence" bson:"sequence"`
}

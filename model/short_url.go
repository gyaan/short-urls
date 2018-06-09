package model

import "gopkg.in/mgo.v2/bson"

type ShortUrl struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	ShortUrl    string        `bson:"short_url" json:"short_url"`
	OriginalUrl string        `bson:"origin_url" json:"original_url"`
	TotalViews  int           `bson:"total_views" json:"total_views"`
}

package dao

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"short-urls/model"
)

type ShortUrlsDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "short_urls"
)

func (m *ShortUrlsDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//get all short urls
func (m *ShortUrlsDao) FindAll() ([]model.ShortUrl, error) {
	var shortUrls [] model.ShortUrl
	err := db.C(COLLECTION).Find(bson.M{}).All(&shortUrls)
	return shortUrls, err
}

//get single short url
func (m *ShortUrlsDao) FindById(id string) (model.ShortUrl, error) {
	var shortUrl model.ShortUrl
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&shortUrl)
	return shortUrl, err
}

//create short url
func (m *ShortUrlsDao) Insert(shortUrl model.ShortUrl) error {
	err := db.C(COLLECTION).Insert(&shortUrl)
	return err
}

//delete short url
func (m *ShortUrlsDao) Delete(shortUrl model.ShortUrl) error {
	err := db.C(COLLECTION).Remove(&shortUrl)
	return err
}

//update short url
func (m *ShortUrlsDao) Update(shortUrl model.ShortUrl) error {
	err := db.C(COLLECTION).UpdateId(shortUrl.Id, &shortUrl)
	return err
}

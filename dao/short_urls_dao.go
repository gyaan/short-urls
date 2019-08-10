package dao

import (
	"github.com/gyaan/short-urls/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ShortUrlsDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	shortUrlCollection = "short_urls"
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
	err := db.C(shortUrlCollection).Find(bson.M{}).All(&shortUrls)
	return shortUrls, err
}

//get single short url
func (m *ShortUrlsDao) FindById(id string) (model.ShortUrl, error) {
	var shortUrl model.ShortUrl
	err := db.C(shortUrlCollection).FindId(bson.ObjectIdHex(id)).One(&shortUrl)
	return shortUrl, err
}

//create short url
func (m *ShortUrlsDao) Insert(shortUrl model.ShortUrl) error {
	err := db.C(shortUrlCollection).Insert(&shortUrl)
	return err
}

//delete short url
func (m *ShortUrlsDao) Delete(shortUrl model.ShortUrl) error {
	err := db.C(shortUrlCollection).Remove(&shortUrl)
	return err
}

//update short url
func (m *ShortUrlsDao) Update(shortUrl model.ShortUrl) error {
	err := db.C(shortUrlCollection).UpdateId(shortUrl.Id, &shortUrl)
	return err
}

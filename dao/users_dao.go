package dao

import (
	"gopkg.in/mgo.v2"
	"log"
	"short-urls/model"
	"gopkg.in/mgo.v2/bson"
)

type UsersDao struct {
	Server   string
	Database string
}

var db1 *mgo.Database

const (
	userCollection = "users"
)

func (m *UsersDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//create one user
func (m *UsersDao) Insert(user model.User) error {
	err := db.C(userCollection).Insert(&user)
	return err
}

//get user details
func (m *UsersDao) FindById(id string) (model.User, error) {
	var user model.User
	err := db.C(userCollection).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

//update user details
func (m *UsersDao) Update(user model.User) error {
	err := db.C(userCollection).UpdateId(user.Id, &user)
	return err
}

//get all users
func (m *UsersDao) FindAll() ([] model.User, error) {

	var users [] model.User
	err := db.C(userCollection).Find(bson.M{}).All(&users)
	return users, err
}

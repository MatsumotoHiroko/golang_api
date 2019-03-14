package dao

import (
	"log"

	. "github.com/MatsumotoHiroko/golang_api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BooksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "books"
)

func (r *BooksDAO) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

func (r *BooksDAO) FindAll() ([]Book, error) {
	var books []Book
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}

func (r *BooksDAO) FindById(id string) (Book, error) {
	var book Book
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

func (r *BooksDAO) Insert(book Book) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}

func (r *BooksDAO) Delete(book Book) error {
	err := db.C(COLLECTION).Remove(&book)
	return err
}

func (r *BooksDAO) Update(book Book) error {
	err := db.C(COLLECTION).UpdateId(book.ID, &book)
	return err
}
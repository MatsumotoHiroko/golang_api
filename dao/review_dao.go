package dao

import (
	"log"

	. "github.com/MatsumotoHiroko/golang_api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReviewsDAO struct {
	Server   string
	Database string
}

var review_db *mgo.Database

const (
	REVIEWS_COLLECTION = "reviews"
)

func (r *ReviewsDAO) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	review_db = session.DB(r.Database)
}

func (r *ReviewsDAO) FindAll() ([]Review, error) {
	var reviews []Review
	err := review_db.C(REVIEWS_COLLECTION).Find(bson.M{}).All(&reviews)
	return reviews, err
}

func (r *ReviewsDAO) FindById(id string) (Review, error) {
	var review Review
	err := review_db.C(REVIEWS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&review)
	return review, err
}

func (r *ReviewsDAO) Insert(review Review) error {
	err := review_db.C(REVIEWS_COLLECTION).Insert(&review)
	return err
}

func (r *ReviewsDAO) Delete(review Review) error {
	err := review_db.C(REVIEWS_COLLECTION).Remove(&review)
	return err
}

func (r *ReviewsDAO) Update(review Review) error {
	err := review_db.C(COLLECTION).UpdateId(review.ID, &review)
	return err
}
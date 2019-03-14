package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Review struct {
	ID bson.ObjectId       `bson:"_id" json:"id"`
	BookId bson.ObjectId `bson:"book_id"`
	Name string           `bson:"name" json:"name" validate:"required,max=50"`
	Comment string           `bson:"comment" json:"comment" validate:"max=100"`
	Value int              `json:"value,string" validate:"required,gte=1,lte=5"`
	CreatedAt time.Time     `bson:"created_at"`
}
package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)
type Book struct {
	ID bson.ObjectId `bson:"_id"`
	Name string      `bson:"name" json:"name" validate:"required,max=50"`
	Price int         `json:"price,string" validate:"required,gte=1,lt=99999999"`
	Difficulty int   `json:"difficulty,string" validate:"required,gte=1,lte=5"`
	Released bool  `json:"released"`
	CreatedAt time.Time     `bson:"created_at"`
}
package persistence

import (
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Cost      int64 //need to tie this with stylist level.
}

type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Address   string
	City      string
	Country   string
	OpenTime  int
	CloseTime int
	Services  []Service
	Stylists  []Stylist
}

type Stylist struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `json:"name"`
	Location string        `json:"location,omitempty"`
	Level    string        `json:"stylistLevel"`
}

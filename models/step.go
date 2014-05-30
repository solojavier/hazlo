package models

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type Step struct {
	Id       bson.ObjectId "_id,omitempty"
	User     string
	Date     time.Time
	Week     int
	Year     int
	Goal     int
	Progress int
}

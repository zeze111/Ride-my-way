package ride

import "gopkg.in/mgo.v2/bson"

type Ride struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	DepartFrom string `bson: "departfrom" json:departfrom,omitempty`
	ArriveAt string `bson: "arriveat" json:arriveat,omitempty`
	Date string `bson: "date" json:date,omitempty`
	Time string `bson: "time" json:time,omitempty`
}

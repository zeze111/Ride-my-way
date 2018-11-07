package model

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RideMyWayDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "rides"
)

func (m *RideMyWayDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *RideMyWayDAO) FindAll() ([]Ride, error) {
	var rides []Ride
	err := db.C(COLLECTION).Find(bson.M{}).All(&rides)
	return rides, err
}

func (m *RideMyWayDAO) FindById(id string) (Ride, error) {
	var ride Ride
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&ride)
	return ride, err
}

func (m *RideMyWayDAO) Insert(ride Ride) error {
	err := db.C(COLLECTION).Insert(&ride)
	return err
}

func (m *RideMyWayDAO) Delete(ride Ride) error {
	err := db.C(COLLECTION).Remove(&ride)
	return err
}

func (m *RideMyWayDAO) Update(ride Ride) error {
	err := db.C(COLLECTION).UpdateId(ride.ID, &ride)
	return err
}

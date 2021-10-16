package persistence

import "gopkg.in/mgo.v2/bson"

// define Booking model
type Booking struct {
	Date    int64
	EventID []byte
	Seats   int
}

// define Event model
type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

// define Location model
type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	City      string
	State     string
	Country   string
	OpenTime  int
	CloseTime int
	Venues    []Venue
}

// define User model
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	First    string
	Last     string
	Age      int
	Bookings []Booking
}

// define Venue model
type Venue struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}

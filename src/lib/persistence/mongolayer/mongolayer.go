package mongolayer

import (
	"goEventsSite/src/lib/persistence"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// define database details
const (
	DB        = "myevents"
	USERS     = "users"
	EVENTS    = "events"
	LOCATIONS = "locations"
)

// define MongoDB session object
type MongoDBLayer struct {
	session *mgo.Session
}

// NewMongoDBLayer initializes the connection to the MongoDB instance
func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	return &MongoDBLayer{
		session: s,
	}, err
}

// getFreshSession creates a fresh database connection to the MongoDB instance
func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}

// AddBooking add a new booking entry to mongodb instance
func (mgoLayer *MongoDBLayer) AddBooking(id []byte, b persistence.Booking) error {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session

	return s.DB(DB).C(USERS).UpdateId(bson.ObjectId(id), bson.M{"$addToSet": bson.M{"bookings": []persistence.Booking{b}}})
}

// AddEvent adds a new event entry to mongodb instance
func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

// 	AddLocation adds a new location entry to the mongodb instance
func (mgoLayer *MongoDBLayer) AddLocation(l persistence.Location) (persistence.Location, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	if !l.ID.Valid() {
		l.ID = bson.NewObjectId()
	}

	return l, s.DB(DB).C(EVENTS).Insert(l)
}

//  AddUser(User) ([]byte, error) adds a new user entry to the mongodb instance
func (mgoLayer *MongoDBLayer) AddUser(u persistence.User) ([]byte, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	if !u.ID.Valid() {
		u.ID = bson.NewObjectId()
	}

	return []byte(u.ID), s.DB(DB).C(EVENTS).Insert(u)
}

// FindAllAvailableEvents gets all events in the mongodb instance
func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	events := []persistence.Event{} // it's a slice because we return multiple events

	return events, s.DB(DB).C(EVENTS).Find(nil).All(&events)
}

// 	FindAllLocations queries the mognodb instance for all available location entries
func (mgoLayer *MongoDBLayer) FindAllLocations() ([]persistence.Location, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	locations := []persistence.Location{}

	return locations, s.DB(DB).C(LOCATIONS).Find(nil).All(&locations)
}

// 	FindAllUserBookings querires the mongodb instance for a given bookingID
func (mgoLayer *MongoDBLayer) FindAllUserBookings(id []byte) ([]persistence.Booking, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	u := persistence.User{}

	return u.Bookings, s.DB(DB).C(USERS).FindId(bson.ObjectId(id)).One(&u)
}

// FindEvent queries mongodb instance for a given eventID
func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	e := persistence.Event{}

	return e, s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
}

// FindEventByName queries mongodb instance for a given event name
func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	e := persistence.Event{}

	return e, s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
}

// 	FindLocation(string) (Location, error)
func (mgoLayer *MongoDBLayer) FindLocation(id string) (persistence.Location, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	l := persistence.Location{}

	return l, s.DB(DB).C(LOCATIONS).Find(bson.M{"_id": bson.ObjectId(id)}).One(&l)
}

// FindUser(string, string) (User, error)
func (mgoLayer *MongoDBLayer) FindUser(f string, l string) (persistence.User, error) {
	s := mgoLayer.getFreshSession() // open the session
	defer s.Close()                 // then close the session - always close open doors

	u := persistence.User{}
	err := s.DB(DB).C(USERS).Find(bson.M{"first": f, "last": l}).One(&u)

	return u, err
}

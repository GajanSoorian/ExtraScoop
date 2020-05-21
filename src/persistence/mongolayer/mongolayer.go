package mongolayer

import (
	"context"
	"fmt"
	"log"

	"github.com/ExtraScoop/src/persistence"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB     = "services"
	USERS  = "users"
	EVENTS = "events"
)

//MongoDBLayer struct has the client object for mongodb driver
type MongoDBLayer struct {
	mongoClient *mongo.Client
}

//NewMongoDBLayer returns a MongoDB Client
func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	clientOptions := options.Client().ApplyURI(connection)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return &MongoDBLayer{
		mongoClient: client,
	}, err
}

func (mongoDBLayer *MongoDBLayer) AddService() {

	fmt.Println("Adding service!!!")
	/*s := mgoLayer.getFreshSession()
	defer s.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e) */
}

/*
func (mgoLayer *MongoDBLayer) FindService(id []byte) (persistence.Service, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Service{}

	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindServiceByName(name string) (persistence.Service, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Service{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableServices() ([]persistence.Service, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	events := []persistence.Service{}
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}
*/
/*
func testCon{
// Set client options
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")
}
*/

package db

import (
	"log"
	"gopkg.in/mgo.v2"
)


func Connect(connectionString string, database string) *mgo.Database {
	log.Println("Trying to connect to mongodb: " + connectionString)
	session, err := mgo.Dial(connectionString)

	if err != nil {
		panic("Can't connect to MongoDB")

	}
	log.Println("MongoDB connection succeed")

	session.SetMode(mgo.Monotonic, true)
	return session.DB(database);
}
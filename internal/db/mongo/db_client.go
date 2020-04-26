/*
 * Copyright Pnoker. All Rights Reserved.
 */

package mongo

import (
	"github.com/google/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

var currentMongoClient MongoClient // Singleton used so that mongoEvent can use it to de-reference readings

type MongoClient struct {
	session  *mgo.Session  // Mongo database session
	database *mgo.Database // Mongo database
}

// Return a pointer to the MongoClient
func NewClient(config Configuration) (MongoClient, error) {
	m := MongoClient{}

	// Create the dial info for the Mongo session
	connectionString := config.Host + ":" + strconv.Itoa(config.Port)
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{connectionString},
		Timeout:  time.Duration(config.Timeout) * time.Millisecond,
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return m, err
	}

	m.session = session
	m.database = session.DB(config.Database)

	currentMongoClient = m // Set the singleton
	return m, nil
}

func (mc MongoClient) CloseSession() {
	if mc.session != nil {
		mc.session.Close()
		mc.session = nil
	}
}

// Get the current Mongo Client
func getCurrentMongoClient() (MongoClient, error) {
	return currentMongoClient, nil
}

// Get a copy of the session
func (mc MongoClient) getSessionCopy() *mgo.Session {
	return mc.session.Copy()
}

func errorMap(err error) error {
	if err == mgo.ErrNotFound {
		err = ErrNotFound
	}
	return err
}

func idToQueryParameters(id string) (name string, value interface{}, err error) {
	if !bson.IsObjectIdHex(id) {
		_, err := uuid.Parse(id)
		if err != nil { // It is some unsupported type of string
			return "", "", ErrInvalidObjectId
		}
		name = "uuid"
		value = id
	} else {
		name = "_id"
		value = bson.ObjectIdHex(id)
	}
	return
}

func idToBsonM(id string) (q bson.M, err error) {
	var name string
	var value interface{}
	name, value, err = idToQueryParameters(id)
	if err != nil {
		return
	}
	q = bson.M{name: value}
	return
}

// Delete from the collection based on ID
func (mc MongoClient) deleteById(col string, id string) error {
	s := mc.getSessionCopy()
	defer s.Close()

	query, err := idToBsonM(id)
	if err != nil {
		return err
	}
	return errorMap(s.DB(mc.database.Name).C(col).Remove(query))
}

func (mc MongoClient) updateId(col string, id string, update interface{}) error {
	s := mc.getSessionCopy()
	defer s.Close()

	query, err := idToBsonM(id)
	if err != nil {
		return err
	}
	return errorMap(s.DB(mc.database.Name).C(col).Update(query, update))
}

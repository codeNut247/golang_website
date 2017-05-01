package database

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jmoiron/sqlx"
	mgo "gopkg.in/mgo.v2"
)

var (
	// BoltDB wrapper (the one we are going to use)
	BoltDB *bolt.DB
	// Mongo wrapper
	Mongo *mgo.Session
	// SQL wrapper
	SQL *sqlx.DB
	// Database info
	databases Info
)

// DbType in order to differentiate with constants
type DbType string

// Info of config.json
type Info struct {
	// Database type
	Type DbType
	// BoltDB info if used
	BoltDB BoltDBInfo
	// MongoDB info if used
	MongoDB MongoDBInfo
}

const (
	// TypeBolt is BoltDB (i didn't catch that)
	TypeBolt DbType = "Bolt"
	// TypeMongoDB is MongoDB (now I get it)
	TypeMongoDB DbType = "MongoDB"
)

// BoltDBInfo contains Path detail for database connection
type BoltDBInfo struct {
	Path string
}

// MongoDBInfo contains details for database connection
type MongoDBInfo struct {
	URL      string
	Database string
}

// Connect to the database
func Connect(dbInfo Info) {
	var err error

	switch dbInfo.Type {
	case TypeBolt:
		// BoltDB is a global variable in package database
		BoltDB, err = bolt.Open(dbInfo.BoltDB.Path, 0600, nil)
		if err != nil {
			log.Println("Bolt Driver Error - ", err)
		}
	case TypeMongoDB:
		// Connect to MongoDB
		Mongo, err = mgo.DialWithTimeout(dbInfo.MongoDB.URL, 5*time.Second)
		if err != nil {
			log.Println("MongoDB Driver Error - ", err)
			return
		}

		// Prevents these errors: read tcp 127.0.0.1:27017: i/o timeout
		Mongo.SetSocketTimeout(1 * time.Second)

		// Check if is alive
		err = Mongo.Ping()
		if err != nil {
			log.Println("Database Error - ", err)
		}

	default:
		log.Println("No registered database in config")
	}
}

// Update modifies Bolt
func Update(bucketName string, key string, dataStruct interface{}) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		// Create the bucket
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}

		// Encode the record
		envodedRecord, err := json.Marshal(dataStruct)
		if err != nil {
			return err
		}

		// Store the record
		err = bucket.Put([]byte(key), envodedRecord)
		if err != nil {
			return err
		}

		// if update successful
		return nil
	})
	return err
}

// View retrieves a record from Bolt
func View(bucketName string, key string, dataStruct interface{}) error {
	err := BoltDB.View(func(tx *bolt.Tx) error {
		// Get the bucket
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		record := bucket.Get([]byte(key))
		if len(record) < 1 {
			return bolt.ErrInvalid
		}

		// Decode the record
		err := json.Unmarshal(record, &dataStruct)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// Delete removes bucket from database
func Delete(bucketName string, key string) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		// Deletes the bucket
		return bucket.Delete([]byte(key))
	})
	return err
}

// CheckConnection returns true if MongoDB is available
func CheckConnection() bool {
	if Mongo == nil {
		Connect(databases)
	} else if Mongo != nil {
		return true
	}
	return false
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}

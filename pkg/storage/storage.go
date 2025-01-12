package storage

import (
	"log"

	"github.com/ithinkiborkedit/GUSH.git/pkg/models"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func InitDB(path string) {
	var err error

	db, err = bbolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open BoltDB: %v", err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Users"))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte("Rooms"))
		return err
	})

	if err != nil {
		log.Fatalf("failed to create buckets: %v", err)
	}

	log.Println("BoltDB initialized")
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func SaveUser(user *models.User) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		return b.Put([]byte(user.Username), []byte(user.Username))
	})
}

func GetUser(username string) (string, error) {
	var user string
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		v := b.Get([]byte(username))
		if v == nil {
			return nil
		}

		user = string(v)
		return nil

	})

	return user, err
}

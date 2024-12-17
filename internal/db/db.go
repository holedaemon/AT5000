package db

import "go.etcd.io/bbolt"

var (
	globalMessagesKey  = []byte("at-global-messages")
	globalResponsesKey = []byte("at-global-responses")
	globalNumbersKey   = []byte("at-numbers")
)

type DB struct {
	f *bbolt.DB
}

func Open(path string) (*DB, error) {
	db := new(DB)
	f, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	db.f = f
	return db, nil
}

func (db *DB) init() error {
	return db.f.Update(func(tx *bbolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(globalMessagesKey); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists(globalResponsesKey); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists(globalNumbersKey); err != nil {
			return err
		}

		return nil
	})
}

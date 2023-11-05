package storage

import (
	"github.com/therobertcrocker/wayfinder/internal/util"
	"go.etcd.io/bbolt"
)

type StorageManager interface {
	// Get retrieves a value from the database.
	Get(bucket string, key []byte) ([]byte, error)

	// Put stores a value in the database.
	Put(bucket string, key []byte, value []byte) error

	// BulkAdd adds a batch of key/value pairs to the database.
	BulkAdd(bucket string, data map[string][]byte) error

	// Close closes the connection to the database.
	Close() error
}

var _ StorageManager = (*BboltStorageManager)(nil)

// BboltConfig holds config values for the database.
type BboltConfig struct {
	buckets []string
}

// BboltStorageManager is a storage manager that uses bbolt.
type BboltStorageManager struct {
	Config BboltConfig
	DB     *bbolt.DB
}

// NewBboltStorageManager creates a new bbolt storage manager.
func NewBboltStorageManager(path string, buckets []string) *BboltStorageManager {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		util.Log.Fatalf("failed to open database: %v", err)
	}

	storage := &BboltStorageManager{
		Config: BboltConfig{
			buckets: buckets,
		},
		DB: db,
	}

	if err := storage.init(); err != nil {
		util.Log.Fatalf("failed to initialize database: %v", err)
	}

	return storage

}

// Get retrieves a value from the database.
func (s *BboltStorageManager) Get(bucket string, key []byte) ([]byte, error) {
	var value []byte
	err := s.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		value = b.Get(key)
		return nil
	})
	return value, err
}

// Put stores a value in the database.
func (s *BboltStorageManager) Put(bucket string, key []byte, value []byte) error {
	return s.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		return b.Put(key, value)
	})
}

// BulkAdd adds a batch of key/value pairs to the database.
func (s *BboltStorageManager) BulkAdd(bucket string, data map[string][]byte) error {
	return s.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		for k, v := range data {
			if err := b.Put([]byte(k), v); err != nil {
				return err
			}
		}
		return nil
	})
}

// Close closes the connection to the database.
func (s *BboltStorageManager) Close() error {
	return s.DB.Close()
}

// init initializes the database.
func (s *BboltStorageManager) init() error {
	return s.DB.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range s.Config.buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

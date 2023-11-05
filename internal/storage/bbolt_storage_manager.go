package storage

import (
	"github.com/therobertcrocker/wayfinder/internal/util"
	"go.etcd.io/bbolt"
)

type StorageManager interface {
	// Close closes the connection to the database.
	Close() error
}

var _ StorageManager = (*BboltStorageManager)(nil)

type BboltConfig struct {
	buckets []string
}

type BboltStorageManager struct {
	Config BboltConfig
	DB     *bbolt.DB
}

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

func (s *BboltStorageManager) Close() error {
	return s.DB.Close()
}

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

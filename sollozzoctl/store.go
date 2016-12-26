package sollozzoctl

import (
	"time"

	"github.com/boltdb/bolt"
	"encoding/json"
)

var (
	projectBucket = []byte("projects")
)

type Store struct {
	path string
	db   *bolt.DB
}

// NewStore returns a new instance of Store.
func NewStore(path string) *Store {
	return &Store{
		path: path,
	}
}

// Path returns the data path.
func (s *Store) Path() string {
	return s.path
}

func (s *Store) Open() error {
	// Open underlying data store.
	db, err := bolt.Open(s.path, 0755, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	s.db = db

	// Initialize all the required buckets.
	if err := s.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(projectBucket)
		return nil
	}); err != nil {
		s.Close()
		return err
	}

	return nil
}

// Close closes the store.
func (s *Store) Close() error {
	if s.db != nil {
		s.db.Close()
	}
	return nil
}

// Ping connects to the database. Returns nil if successful.
func (s *Store) Ping() error {
	return s.db.View(func(tx *bolt.Tx) error {
		return nil
	})
}

func (s *Store) Put(key []byte, content []byte) {
	store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectBucket)

		bucket.Put(key, content)

		return nil;
	})
}

func (s *Store) Delete(key []byte) {
	store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectBucket)

		bucket.Delete(key);

		return nil
	})
}

func (s *Store) Get(key []byte, t interface{}) {
	store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectBucket)

		encoded := bucket.Get(key)

		json.Unmarshal(encoded, t)

		return nil;
	})
}

func (s *Store) forEach(fn func(k, v []byte) error) {
	store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectBucket)

		bucket.ForEach(fn)

		return nil;
	})
}
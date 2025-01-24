package store

import "sync"

type Store struct {
	Codes *sync.Map
	Stats *sync.Map
}

func NewStore() *Store {
	return &Store{Codes: new(sync.Map), Stats: new(sync.Map)}
}

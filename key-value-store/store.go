package key_value_store

import (

)

type Store struct {
	filePath string
}

func NewStore(filePath string) (*Store, error) {
	return &Store{
		filePath: filePath,
	}, nil
}

func (reader *Store) Get(key string) (string, error) {
	return "", nil
}

func (reader *Store) Set(key string, value string) (string, error) {
	return "", nil
}

package store

import (
	"encoding/json"
	"os"
)

type Store struct {
	path string
}

func New(path string) *Store {
	return &Store{
		path: path,
	}
}

// Write function to write data
func (s *Store) Write(data interface{}) {
	f, err := os.Create(s.path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(data)
	if err != nil {
		panic(err)
	}
}

// Read function to read data
func (s *Store) Read() interface{} {
	f, err := os.Open(s.path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var data interface
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}

package store

import (
	"fmt"
	"os"
)

// Write function to write data
func Write(data string, filepath string) (filename string, err error) {
	// need data
	// storage type - file, db, etc.?
	// connectivitity or storage location
	// file storage
	// what would I need for file storage?
	// file name? I could auto generate this
	// where do I want to store?
	// path for file storage

	//Steps
	// validate path exists
	// -- needs error check
	// generate file name
	// create file at path
	// -- need error check
	// write data to file
	// -- need error check

	// Check if path exists
	// what if file already exists?
	// Do we over-write or append?
	// minimum would be over-write
	// could have an append API as well
	// implement empty interface

	filename = filepath + "/data.txt"
	fmt.Println("File is ", filename)
	f, err2 := os.Create(filename)

	if err2 != nil {
		return "file not created", err2
	}

	defer f.Close()

	_, err3 := f.WriteString(data)

	if err3 != nil {
		return "Error writing data to file", err3
	}
	return filename, nil
}

// Read function to read data
func Read(filepath string) (data string, err error) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

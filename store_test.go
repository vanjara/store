package store_test

import (
	"store"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Parallel()
	input := "test data"
	// we could take file path instead of directory path
	dirpath := "/tmp"
	got, err := store.Write(input, dirpath)
	if err != nil {
		t.Errorf("error")
	}
	want := "/tmp/data.txt"
	if want != got {
		t.Fatalf("wanted %v, but got %v", want, got)
	}
	// needs to validate it wrote anything to the file
	// how could we check that
	// we just read it back and validate we wrote correctly
	// should add failure check too
}

func TestFailWrite(t *testing.T) {
	t.Parallel()
	input := "test fail data"
	var errorExpected bool
	// we could take file path instead of directory path
	dirpath := "/tmp1"
	got, err := store.Write(input, dirpath)
	// if err != nil {
	// 	t.Errorf("error %v", err)
	// }
	want := "/tmp1/data.txt"
	errorExpected = true
	if errorExpected != (err != nil) {
		t.Fatalf("Given input %q & dirpath %q, unexpected error Status: %v", input, dirpath, err)
	}
	if !errorExpected && want != got {
		t.Errorf("Given input: %q, want %q, got %q\n", input, want, got)
	}

	// if want != got {
	// 	t.Fatalf("wanted %v, but got %v", want, got)
	// }
	// needs to validate it wrote anything to the file
	// how could we check that
	// we just read it back and validate we wrote correctly
	// should add failure check too
}

func TestRead(t *testing.T) {
	t.Parallel()
	filepath := "/tmp/data.txt"
	got, err := store.Read(filepath)
	if err != nil {
		t.Errorf("error")
	}
	want := "test data"
	if want != got {
		t.Fatalf("wanted %v, but got %v", want, got)
	}
}

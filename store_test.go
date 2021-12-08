package store_test

import (
	"store"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWriteRead(t *testing.T) {
	t.Parallel()
	want := struct{ Name string }{Name: "John"}
	// we could take file path instead of directory path
	dirpath := "test.dat"
	s := store.New(dirpath)
	s.Write(want)
	s2 := store.New(dirpath)
	got := s2.Read()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

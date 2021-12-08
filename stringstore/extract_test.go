package stringstore_test

import (
	"store/stringstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractContentsMapStore(t *testing.T) {
	t.Parallel()
	s := stringstore.NewMapStore()
	want := "Hello, world"
	s.Put(want)
	got := stringstore.ExtractContents(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestExtractContentsStructStore(t *testing.T) {
	t.Parallel()
	s := stringstore.NewStructStore()
	want := "Hello, world"
	s.Put(want)
	got := stringstore.ExtractContents(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

type ArbitraryStore struct{}

func (a ArbitraryStore) Put(string) {}

func (a ArbitraryStore) Get() string { return "" }

func TestExtractContentsArbitraryStore(t *testing.T) {
	t.Parallel()
	s := ArbitraryStore{}
	want := "Hello, world"
	s.Put(want)
	got := stringstore.ExtractContents(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

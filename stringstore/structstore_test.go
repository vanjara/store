package stringstore_test

import (
	"store/stringstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPutGet(t *testing.T) {
	t.Parallel()
	s := stringstore.NewStructStore()
	want := "Hello, world"
	s.Put(want)
	got := s.Get()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}


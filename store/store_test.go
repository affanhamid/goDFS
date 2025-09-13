package store

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedFileName := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s, want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.FileName != expectedFileName {
		t.Errorf("have %s, want %s", pathKey.PathName, expectedFileName)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)

	for i := range 100 {
		key := fmt.Sprintf("foo_%d", i)

		data := []byte("some jpg bytes")

		_, err := s.writeStream(key, bytes.NewReader(data))

		assert.NoError(t, err)

		assert.True(t, s.Has(key))

		n, r, err := s.Read(key)
		assert.NoError(t, err)
		assert.Equal(t, n, int64(14))

		b, _ := io.ReadAll(r)
		assert.Equal(t, string(b), string(data))

		assert.NoError(t, s.Delete(key))

		fmt.Println("key", key, "has", s.Has(key))
		assert.False(t, s.Has(key))

	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	assert.NoError(t, s.Clear())
}

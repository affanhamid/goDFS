package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: NOPPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader(([]byte("some jpg bytes")))

	s.writeStream("mySpecialPicture", data)
}

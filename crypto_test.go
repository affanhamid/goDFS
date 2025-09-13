package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyEncryptDecrypt(t *testing.T) {
	payload := "foo not bar"
	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)
	key := NewEncryptionKey()

	_, err := copyEncrypt(key, src, dst)
	assert.NoError(t, err)
	fmt.Println(dst.String())

	out := new(bytes.Buffer)
	nw, err := copyDecrypt(key, dst, out)
	assert.NoError(t, err)
	assert.Equal(t, out.String(), payload)
	assert.Equal(t, nw, 16+len(payload))
}

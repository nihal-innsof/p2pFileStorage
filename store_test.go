package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {}

func TestStore(t *testing.T) {
	storeOpts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}
	store := NewStore(storeOpts)

	data := bytes.NewReader([]byte("Test data"))
	if err := store.writeStream("meow", data); err != nil {
		assert.Nil(t, err)
	}
}

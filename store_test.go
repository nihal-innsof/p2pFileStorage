package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "meowmeow"
	pathName := CASPathTransformFunc(key)
	expectedPathName := "b6ccb/4ece5/454dc/ae517/78b3e/239eb"
	if pathName != expectedPathName {
		t.Errorf("Expected %s, got %s", expectedPathName, pathName)
	}
}

func TestStore(t *testing.T) {
	storeOpts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	store := NewStore(storeOpts)

	data := bytes.NewReader([]byte("Test data"))
	if err := store.writeStream("meow", data); err != nil {
		assert.Nil(t, err)
	}
}

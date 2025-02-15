package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMeAdd(t *testing.T) {
	assert.Equal(t, 3, testMeAdd(1, 2))
}

func TestMeSub(t *testing.T) {
	assert.Equal(t, 3, testMeSub(1, 2))
}

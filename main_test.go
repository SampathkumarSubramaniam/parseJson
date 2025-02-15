package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMe(t *testing.T) {
	assert.Equal(t, 3, testMe(1, 2))
}

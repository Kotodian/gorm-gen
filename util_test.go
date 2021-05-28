package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstLower(t *testing.T) {
	str := "Test"
	lower := firstLower(str)
	assert.Equal(t, "test", lower)
}

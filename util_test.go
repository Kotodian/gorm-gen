package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsGorm(t *testing.T) {
	ok := isGorm("name", `gorm:"column:name"`)
	assert.Equal(t, true, ok)
}

func TestFirstLower(t *testing.T) {
	str := "Test"
	lower := firstLower(str)
	assert.Equal(t, "test", lower)
}

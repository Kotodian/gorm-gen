package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValue(t *testing.T) {
	key, value := keyValue("json:\"name\"")
	t.Log(key)
	t.Log(value)
}

func TestExtractTag(t *testing.T) {
	tags := extractTag("`json:\"name\" gorm:\"column:name\" create:\"scope:10\"`")
	assert.Equal(t, "name", tags[1].Value("column"))
}

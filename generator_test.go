package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerator_ParseDir(t *testing.T) {
	g := &Generator{
		files: make([]string, 0),
	}
	dir := "/Users/linqiankai/go/src/gorm-gen/examples"
	g.ParseDir(dir)
	fmt.Println(g.files)
	pkgName := dir[strings.LastIndex(dir, "/")+1:]
	fmt.Println(pkgName)
	assert.NotNil(t, g.files)
}

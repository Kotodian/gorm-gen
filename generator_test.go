package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/packages"
	"strings"
	"testing"
)

func TestGenerator_ParseDir(t *testing.T) {
	g := &Generator{
		files: make([]string, 0),
	}
	dir := "/Users/linqiankai/go/src/gorm-gen/examples"
	g.parseDir(dir)
	fmt.Println(g.files)
	pkgName := dir[strings.LastIndex(dir, "/")+1:]
	fmt.Println(pkgName)
	assert.NotNil(t, g.files)
}

func TestParseFile(t *testing.T) {
	g := &Generator{
		files: make([]string, 0),
	}
	dir := "/Users/linqiankai/go/src/gorm-gen/examples/models/common"
	g.parseDir(dir)
	files, err := g.parseFile()
	assert.Nil(t, err)
	for _, file := range files {
		t.Log(file.pkg)
		t.Log(file.structs)
	}
}

func TestGenerate(t *testing.T) {
	g := &Generator{
		output: "/Users/linqiankai/go/src/gorm-gen/examples/dao",
		files:  make([]string, 0),
	}
	g.generate("/Users/linqiankai/go/src/gorm-gen/examples/models")
}

func TestPackagePath(t *testing.T) {
	load, err := packages.Load(
		&packages.Config{
			Dir:  "/Users/linqiankai/go/src/gorm-gen/examples/models/common",
			Mode: packages.NeedModule,
		})
	assert.Nil(t, err)
	t.Log(load[0].ID)
}

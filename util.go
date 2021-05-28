package main

import (
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path"
	"strings"
)

func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func isGorm(field string, tag string) bool {
	childTag := strings.Split(tag, ":")[1]
	childTag = strings.TrimPrefix(childTag, `"`)
	if childTag != "column" {
		return true
	}
	tagField := strings.Split(tag, ":")[2]
	tagField = strings.TrimSuffix(tagField, `"`)
	return field == tagField
}

func filename(file string) (string, string) {
	dir, file := path.Split(file)
	return dir, file
}

func firstLower(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func pkgPath(str string) string {
	load, err := packages.Load(&packages.Config{
		Dir:  str,
		Mode: packages.NeedModule,
	})
	if err != nil {
		panic(err)
	}
	return load[0].ID
}

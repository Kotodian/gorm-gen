package main

import (
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

func filename(file string) string {
	_, file = path.Split(file)
	return strings.TrimSuffix(file, ".go")
}

func firstLower(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

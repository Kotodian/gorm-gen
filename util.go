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

func isGorm(field string, tags []*Tag) bool {
	for _, t := range tags {
		if t.name == "gorm" {
			if t.value == "" {
				return t.child.value == field
			} else {
				return t.value == field
			}
		}
	}
	return false
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

func tag(str string) string {
	return ""
}

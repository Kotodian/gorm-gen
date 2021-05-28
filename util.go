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

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

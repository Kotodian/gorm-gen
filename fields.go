package main

import (
	"strings"
)

// Field 字段
type Field struct {
	// 字段名称
	name string
	// 字段类型
	goType string
	// tag
	tag string
	// tags
	tags Tags
}

type Tags []*Tag

func (tags Tags) Exists(key string) *Tag {
	for _, t := range tags {
		if t.name == key {
			return t
		} else {
			if t.child.name == key {
				return t
			}
		}
	}
	return nil
}
func (tags Tags) Gorm(field string) bool {
	if tag := tags.Exists("gorm"); tag != nil {
		return tag.value == firstLower(field) || tag.Value("column") == firstLower(field)
	}
	return false
}

// Tag 一个字段的tag
type Tag struct {
	name  string
	value string
	child *Tag
}

func (t *Tag) Value(key string) string {
	if t.name == key {
		return t.value
	}
	if t.child == nil {
		return ""
	}
	return t.child.Value(key)
}

// extractTag 将形如`json:"name" gorm:"column:name" create:"scope:10"`转换成
// json (value: name)
// gorm
// 	-- column (value: name)
// create
//  -- scope (value: 10)
func extractTag(t string) Tags {
	t = strings.Trim(t, "`")
	tags := make([]*Tag, 0)
	for _, s := range strings.Split(t, " ") {
		// 例如 json:"name"
		// key=json, value=name
		key, value := keyValue(s)
		tag := &Tag{name: key}
		childKey, childValue := keyValue(value)
		if childKey != "" && childValue != "" {
			childTag := &Tag{name: childKey, value: childValue}
			tag.child = childTag
		} else {
			tag.value = value
		}
		tags = append(tags, tag)
	}
	return tags
}

func keyValue(t string) (string, string) {
	index := strings.Index(t, ":")
	if index == -1 {
		return "", ""
	}
	return t[:index], strings.Trim(t[index+1:], "\"")
}

// childTag 子tag 例如 `gorm:"column:field"` column为子tag
// `gorm:"foreignKey:field"` foreignKey为子tag

// String 形如 "name string"
func (field *Field) String() string {
	return firstLower(field.name) + " " + field.goType
}

func (field *Field) Strings() string {
	return firstLower(field.name) + " []*" + field.goType
}

// DB 将go中的字段名称转成数据库中的字段名称
func (field *Field) DB() string {
	return snakeString(field.name)
}

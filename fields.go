package main

import "strings"

// Field 字段
type Field struct {
	// 字段名称
	name string
	// 字段类型
	goType string
	// tag
	tag string
}

func NewField(name string, typ string, tag string) *Field {
	return &Field{name, typ, tag}
}

// String 形如 "name string"
func (field *Field) String() string {
	return field.name + " " + field.goType
}

// childTag 子tag 例如 `gorm:"column:field"` column为子tag
// `gorm:"foreignKey:field"` foreignKey为子tag
func (field *Field) childTag() string {
	temp := strings.Trim(field.tag, "\"")
	childTag := strings.Split(temp, ":")[2]
	return childTag
}

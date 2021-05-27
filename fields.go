package main

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

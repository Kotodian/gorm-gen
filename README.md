# Gorm Generator
声明: 该项目为根据结构体生成orm操作并只支持go 1.16版本以后其次请使用gorm v2并且只支持go mod

下载项目:

`go get github.com/Kotodian/gorm-gen`

使用:

`gorm-gen --dir=<项目结构体路径> --output=<生成文件路径>`

结构体example:

```go
type User struct {
	Name    string `gorm:"column:name" query:"equal,scope"`
	Age     int    `gorm:"column:age" query:"scope"`
}
```

结构体中的字段的tag中必须包含gorm,其次,query代表生成新的查询,equal代表Where("field = ?", xxx), scope代表范围查询 >=, <=

TODO:

​	关联结构体生成orm
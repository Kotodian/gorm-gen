package common

type User struct {
	Name    string `gorm:"column:name" query:"equal,scope"`
	Age     int    `gorm:"column:age" query:"scope"`
	Company Company
}

type Company struct {
	Name string `gorm:"column:name"`
}

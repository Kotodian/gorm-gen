package common

type User struct {
	Name    string `gorm:"column:name" create:"scope:10"`
	Age     int    `gorm:"column:age" query:"scope:10,20"`
	Company Company
}

type Company struct {
	Name string `gorm:"column:name"`
}

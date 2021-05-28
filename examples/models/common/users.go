package common

type User struct {
	Name    string `gorm:"column:name"`
	Age     int    `gorm:"column:age"`
	Company Company
}

type Company struct {
	Name string `gorm:"column:name"`
}

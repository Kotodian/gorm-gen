package common

type Status int
type User struct {
	Name    string `gorm:"column:name" query:"equal,scope"`
	Age     int    `gorm:"column:age" query:"scope"`
	Company Company
	Status  Status `gorm:"column:status" query:"equal"`
}

type Company struct {
	Name string `gorm:"column:name"`
}

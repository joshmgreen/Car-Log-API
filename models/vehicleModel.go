package models

type Vehicle struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Year    int    `json:"year"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Mileage int    `json:"mileage"`
}

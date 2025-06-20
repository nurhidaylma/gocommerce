package domain

type Address struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"index"`
	Province   string
	City       string
	District   string
	PostalCode string
	Detail     string
}

package domain

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Password string
	Role     string
	Store    Store
}

type Store struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	UserID uint
}

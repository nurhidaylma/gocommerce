package domain

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	CategoryID  uint
	UserID      uint
	Price       float64
	Stock       int
	ImageURL    string
	Description string
}

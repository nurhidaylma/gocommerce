package domain

type Store struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"uniqueIndex"`
	Name   string
	Logo   string
}

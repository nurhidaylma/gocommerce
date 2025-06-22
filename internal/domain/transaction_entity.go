package domain

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	AddressID uint
	Total     int
	Items     []TransactionItem `gorm:"foreignKey:TransactionID"`
	Status    TransactionStatus
}

type TransactionStatus int

const (
	Pending   TransactionStatus = 1
	Paid      TransactionStatus = 2
	Cancelled TransactionStatus = -1
)

type TransactionItem struct {
	ID            uint
	TransactionID uint
	ProductID     uint
	Quantity      int
	Price         int
	LogProduct    LogProduct `gorm:"foreignKey:ItemID"`
}

type LogProduct struct {
	ID          uint `gorm:"primaryKey"`
	ItemID      uint // FK TransactionItem
	Name        string
	Description string
	Price       int
	ImageURL    string
}

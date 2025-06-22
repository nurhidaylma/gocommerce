package dto

type CreateTransactionInput struct {
	AddressID uint                    `json:"address_id"`
	Items     []CreateTransactionItem `json:"items"`
}

type CreateTransactionItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

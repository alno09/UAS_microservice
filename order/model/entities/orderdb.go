package entities

type Order struct {
	CustomerId uint  `json:"customer_id"`
	CatalogId  uint  `json:"catalog_id"`
	Amount     int64 `json:"amount"`
	Price      int64 `json:"price"`
}

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Customer struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

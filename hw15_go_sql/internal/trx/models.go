package trx

type OrderFullParams struct {
	UserID   string    `json:"userId"`
	Products []Product `json:"products"`
}

type Product struct {
	ID    int64 `json:"id"`
	Count int64 `json:"count"`
}

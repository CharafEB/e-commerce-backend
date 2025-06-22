package dots

type OrdersData struct {
	OrderID      string
	ProductsID   string
	UserFullName string
	UserState    string
	UserCity     string
	FullAddress  string
	PhoneNUmber  string
	TotalPrice   float64
	OrderItem    string
	Size         []string
	Color        []string
}

type Orders struct {
	OrderID      string      `json:"order_id"`     // VARCAHR(26) ULID
	UserFullName string      `json:"full_name"`    // TEXT
	UserState    string      `json:"user_state"`   // TEXT
	UserCity     string      `json:"user_city"`    // TEXT
	FullAddress  string      `json:"full_address"` // TEXT
	PhoneNUmber  string      `json:"phone_number"` // VARCHAR(10)
	TotalPrice   float64     `json:"total_price"`  // MUMERIQ(6,3)
	OrderItem    []OrderItem `json:"orderitem"`    // no refrece in the table
}

type OrderItem struct {
	OrderItem string   `json:"orderitem_id"` // VARCAHR(26) ULID
	ProductID string   `json:"product_id"`   // VARCAHR(26) ULID
	OrderID   string   `json:"order_id"`     // VARCAHR(26) ULID
	Quantity  int      `json:"quantity"`
	Sizes     []string `json:"sizes_omitempty"`  // no refrece in the table
	Colors    []string `json:"colors_omitempty"` // no refrece in the table
}

type SizeOrderItem struct {
	OrderItemID string `json:"orderitem_id"` // VARCAHR(26) ULID
	SizeValue   string `json:"size_value"`
}

type ColorOrderItem struct {
	ColoreValue string `json:"color_value"`
	OrderItemID string `json:"orderitem_id"` // VARCAHR(26) ULID
}

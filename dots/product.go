package dots

type Product struct {
	ProductsID          string
	ProductsName        string
	ProductsShortDes    string
	ProductsDescription string
	Price               float64
	ImgURl              string
	ImgesURL            []string
	Size                []string
	Color               []string
	Categories          string
}

type Products struct {
	ProductsID          string  `json:"product_id"` // VARCAHR(26) ULID
	ProductsName        string  `json:"prod_name"`
	ProductsShortDes    string  `json:"prod_min_dec"`
	ProductsDescription string  `json:"prod_doc"`
	Price               float64 `json:"prod_price"`
	ImgURl              string  `json:"main_img"`
	Categories          string  `json:"categories"`
}

type ProductsSize struct {
	SizeValue string `json:"size_value"`
	ProdID    string `json:"product_id"` // VARCAHR(26) ULID
}

type ProductsColor struct {
	ColorValue string `json:"color_value"`
	ProdID     string `json:"product_id"` // VARCAHR(26) ULID
}

type ProductsImges struct {
	ImgID    int    `json:"img_id"`
	ProdID   string `json:"product_id"` // VARCAHR(26) ULID
	ImgValue string `json:"img_value"`
}

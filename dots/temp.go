package dots

import (
	"context"

	"github.com/google/uuid"
)

type ArticleGetter interface {
	GetProducts(ctx context.Context) ([]Products, error)
}

type Indexer interface {
	IndexArticles(ctx context.Context) error
}

type SearchResult struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ShortDes string `json:"short_description"`
	Class    string `json:"class"`
	ImgURl   string `json:"image_url"`
}



type Order struct {
	OrderId          uuid.UUID `json:"order_id"`
	ProdID           uuid.UUID `json:"prod_id"`
	ProductID        string
	UserName         string `json:"user_fullname"`
	UserNumber       string `json:"user_phnumeber"`
	UserState        string `json:"user_state "`
	UserMunicipality string `json:"user_municipality"`
	UserAddress      string `json:"user_address"`
}

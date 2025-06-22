package model

import (
	"context"
	"database/sql"
	"github/think.com/dots"

	"github/think.com/migler/index"

	"github.com/blevesearch/bleve"
)

type Prod struct {
	DB *sql.DB
}

type Store struct {
	Update interface {
		UpdateEvent(ctx context.Context, queryVal []string) error
		Modifierevent(ctx context.Context, queryVal []string) error
	}

	Get interface {
		GetContentSearch(ctx context.Context, val string) ([]dots.Products, error)

		// GetProducts retrieves all products from the database.
		GetProductContent(ctx context.Context, productID string) (*dots.Product, error)

		// GetByCategorie retrieves products by category.
		GetByCategorie(ctx context.Context, category string) ([]dots.Products, error)

		//GetByPriceUnder retrieves products with a price under a specified value.
		GetByPriceUnder(ctx context.Context, priceUnder float64) ([]dots.Products, error)

		// GetByPriceUper retrieves products with a price over a specified value.
		GetByPriceUper(ctx context.Context, priceUper float64) ([]dots.Products, error)

		// GetProducts retrieves all products .
		GetProducts(ctx context.Context) ([]dots.Products, error)

		// GetCategories retrieves all Categories.
		GetCategories(ctx context.Context) ([]dots.Categories, error)

		// GetColors retrieves all colors.
		GetColors(ctx context.Context) ([]dots.Color, error)
		// GetSizes retrieves all sizes.
		GetSizes(ctx context.Context) ([]dots.Size, error)
	}

	Post interface {
		// CreateProducts creates a new product and indexes it.
		CreateProducts(ctx context.Context, prod *dots.Product, Res bleve.Index) error

		//CreateOrder create a new order for user
		CreateOrder(ctx context.Context, order *dots.Orders, Res bleve.Index) error

		// CreateSize creates a new Size for a product.
		AddSize(ctx context.Context, size string) error
		// CreateColor creates a new Color for a product.
		AddColor(ctx context.Context, color string) error
		// CreateImges creates a new Categories for a product.
		AddCategories(ctx context.Context, categories string) error

		// CreateImges creates a new Imges for a product.
		AddImges(ctx context.Context, imgs []string, prodID string) error
	}
	Delete interface {
		DeleteProducts(ctx context.Context, ProductID string) error
		DeleteSize(ctx context.Context, SizeValue string) error
		DeleteColor(ctx context.Context, ColorValue string) error
		DeleteCategories(ctx context.Context, CategorieValue string) error
	}
	Orders interface {
		GetOrders(ctx context.Context) ([]dots.Orders, error)
		GetOrderByID(ctx context.Context, orderID string) (*dots.Orders, error)
	}

	GetProducts dots.ArticleGetter

	Indexer interface {
		IndexArticles(ctx context.Context) error
	}
	Index dots.Indexer
}

func NewStore(db *sql.DB) Store {
	if db == nil {
		panic("nil pointer passed to NewStore")
	}
	articleRepo := &Prod{DB: db}

	return Store{
		Update:      articleRepo,
		Delete:      articleRepo,
		Orders:      articleRepo,
		Post:        articleRepo,
		GetProducts: articleRepo,
		Get:         articleRepo,
		Indexer:     index.NewIndexService(articleRepo),
	}
}

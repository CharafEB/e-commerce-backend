package index

import (
	"github/think.com/dots"

	"github.com/blevesearch/bleve"
)

type BelveIndex struct {
	BelveHandler  BelveHandler
	BelveIndexing BelveIndexing
	BelveSerch    BelveSerch
}

type BelveHandler interface {
	DBHandlerIndex()
}

type BelveIndexing interface {
	AssNewIndex(prod dots.Products) error
}

type BelveSerch interface {
	BSearch(val string, Res bleve.Index) (map[string]interface{}, error)
}

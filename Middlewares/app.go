package middlewares

import (
	"github/think.com/model"

	"github.com/blevesearch/bleve"
)

type Application struct {
	Address          string
	Storge           model.Store
	BleveSearchIndex bleve.Index
	Test             string
}


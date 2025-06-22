package index

import (
	"log"

	"github.com/blevesearch/bleve"
)
	
func indexing() {
	index, err := bleve.New("articel.bleve", bleve.NewIndexMapping())

	if err != nil {
		log.Fatal(err)
	}

	doc := map[string]interface{}{
		"title":   "Introduction to Bleve",
		"content": "Bleve is a full-text search library for Go.",
	}

	err = index.Index("1" , doc)
	if err != nil {
		log.Fatal(err)
	}


	log.Println("we add new index to the index file")
}

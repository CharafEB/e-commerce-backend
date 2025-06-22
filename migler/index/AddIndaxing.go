package index

import (
	"fmt"
	"log"

	"github/think.com/dots"

	"github.com/blevesearch/bleve"
)

func AssNewIndex(prod dots.Products, Res bleve.Index) error {
	log.Println("Adding a new index to the index file")
	fmt.Println("Adding a new index to the index file")
	fmt.Print(prod)

	doc := map[string]interface{}{
		"title":             prod.ProductsName,
		"short_description": prod.ProductsShortDes,
		"image_url":         prod.ImgURl,
		"class":             prod.Categories,
	}

	fmt.Println("The title is:", doc)
	if err := Res.Index(prod.ProductsID, doc); err != nil {
		log.Print("Failed to index document:", err)
		return err // Exit if indexing fails
	}

	log.Println("Successfully added a new index to the index file")
	return nil
}

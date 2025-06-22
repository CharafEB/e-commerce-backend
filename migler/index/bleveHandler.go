package index

import (
	"context"
	"log"

	"github/think.com/dots"

	"github.com/blevesearch/bleve"
)

type IndexService struct {
	articleRepo dots.ArticleGetter
}

func NewIndexService(articleRepo dots.ArticleGetter) *IndexService {
	return &IndexService{articleRepo: articleRepo}
}

func (s *IndexService) IndexArticles(ctx context.Context) error {
	index, err := bleve.New("Searchindex.bleve", bleve.NewIndexMapping())
	if err != nil {
		return err
	}
	defer index.Close()
	products, err := s.articleRepo.GetProducts(ctx)
	if err != nil {
		return err
	}

	for _, product := range products {
		doc := map[string]interface{}{
			"title":             product.ProductsName,
			"short_description": product.ProductsShortDes,
			"image_url":         product.ImgURl,
			"class":             product.Categories,
		}
		if err := index.Index(product.ProductsID, doc); err != nil {
			return err
		}

		// log.Printf("Indexed article with ID: %s The title is : %s The cintent is : %s the IMGURL : %s the Class is : %s\n", article.ID , article.Title , article.ShortDes , article.ImgURl , article.Class)
	}

	log.Println("Indexing has been done")
	return nil
}

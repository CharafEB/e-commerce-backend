package model

import (
	"context"
	"fmt"
	"log"

	"github.com/oklog/ulid/v2"
)

// --> CreateSize creates a new Size for a product.
func (artC *Prod) addProductsSize(ctx context.Context, sizes []string, prodID string) error {

	if prodID == "" {
		return nil
	}

	if len(sizes) == 0 {
		return nil

	}

	query := `INSERT INTO ProductsSize (size_value, product_id) VALUES ($1, $2)`

	for _, size := range sizes {

		if size == "" {
			continue
		}
		_, err := artC.DB.ExecContext(ctx, query, size, prodID)
		if err != nil {
			log.Printf("Filesd to sert: %s for the product_id:%s || %v \n", size, prodID, err)
			return fmt.Errorf("Error adding size : %w", err)
		}
	}

	return nil
}

// --> CreateColor creates a new Color for a product.
func (artC *Prod) addProductsColor(ctx context.Context, colors []string, prodID string) error {

	if prodID == "" {
		return nil
	}

	if len(colors) == 0 {
		return nil

	}

	query := `INSERT INTO ProductsColor (color_value, product_id) VALUES ($1, $2)`

	for _, color := range colors {

		if color == "" {
			continue
		}
		_, err := artC.DB.ExecContext(ctx, query, color, prodID)
		if err != nil {
			log.Printf("Filesd to sert: %s for the product_id:%s || %v \n", color, prodID, err)
			return fmt.Errorf("Error adding color : %w", err)
		}
	}

	return nil
}

// --> CreateImges creates a new Imges for a product.
func (artC *Prod) addProductsImges(ctx context.Context, imgs []string, prodID string) error {

	if prodID == "" {
		return nil
	}

	if len(imgs) == 0 {
		return nil

	}

	query := `INSERT INTO ProductsImges (product_id , img_value) VALUES ($1, $2)`

	for _, img := range imgs {

		if img == "" {
			continue
		}
		_, err := artC.DB.ExecContext(ctx, query, prodID, img)
		if err != nil {
			log.Printf("Filesd to sert: %s for the product_id:%s || %v \n", img, prodID, err)
			return fmt.Errorf("Error adding img : %w", err)
		}
	}

	return nil
}

// --> CreateImges creates a new Imges for a product.
func (artC *Prod) generateuild() string {
	return ulid.Make().String()
}

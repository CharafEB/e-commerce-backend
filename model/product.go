package model

import (
	"context"
	"fmt"
	"log"
)

// AddSize : will hande addin a new size in size table
func (artC *Prod) AddSize(ctx context.Context, size string) error {

	if size == " " {
		return nil

	}

	query := `INSERT INTO Size (size_value) VALUES ($1)`

	_, err := artC.DB.ExecContext(ctx, query, &size)
	if err != nil {
		return fmt.Errorf("Error adding size : %w", err)
	}

	return nil
}

// AddColor : will hande addin a new color in color table
func (artC *Prod) AddColor(ctx context.Context, color string) error {

	if color == " " {
		return nil

	}

	query := `INSERT INTO Color (color_value) VALUES ($1)`

	_, err := artC.DB.ExecContext(ctx, query, &color)
	if err != nil {
		return fmt.Errorf("Error adding color : %w", err)
	}

	return nil
}

// AddCategories : will hande addin a new Categorie in color Categories
func (artC *Prod) AddCategories(ctx context.Context, categories string) error {

	if categories == " " {
		return fmt.Errorf("categorie_value is empty")

	}

	query := `INSERT INTO Categories (categorie_value) VALUES ($1)`
	_, err := artC.DB.ExecContext(ctx, query, categories)
	if err != nil {
		return fmt.Errorf("Error adding Categorie : %w", err)
	}

	return nil
}

// --> CreateImges creates a new Imges for a product.
func (artC *Prod) AddImges(ctx context.Context, imgs []string, prodID string) error {

	if prodID == "" {
		return nil
	}

	if len(imgs) == 0 {
		return nil

	}

	query := `INSERT INTO imgs (prod_id, img_value) VALUES ($1, $2)`

	for _, img := range imgs {

		if img == "" {
			continue
		}
		_, err := artC.DB.ExecContext(ctx, query, prodID, img)
		if err != nil {
			log.Printf("Filesd to sert: %s for the prod_id:%s || %v \n", img, prodID, err)
			return fmt.Errorf("Error adding img : %w", err)
		}
	}

	return nil
}

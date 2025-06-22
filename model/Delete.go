package model

import (
	"context"
	"fmt"
)

func (artC *Prod) DeleteProducts(ctx context.Context, ProductID string) error {
	if ProductID == "" {
		return fmt.Errorf("ProductID is empty")
	}
	query := `DELETE FROM Products WHERE product_id = $1`
	row, err := artC.DB.Exec(query, ProductID)
	if err != nil {
		return err
	}
	fmt.Print(row)
	return nil
}

func (artC *Prod) DeleteSize(ctx context.Context, SizeValue string) error {
	if SizeValue == "" {
		return fmt.Errorf("SizeValue is empty")
	}
	query := `DELETE FROM Size WHERE size_value = $1`
	row, err := artC.DB.Exec(query, SizeValue)
	if err != nil {
		return err
	}
	fmt.Print(row)
	return nil
}

func (artC *Prod) DeleteColor(ctx context.Context, ColorValue string) error {
	if ColorValue == "" {
		return fmt.Errorf("ColorValue is empty")
	}
	query := `DELETE FROM Color WHERE color_value = $1`
	row, err := artC.DB.Exec(query, ColorValue)
	if err != nil {
		return err
	}
	fmt.Print(row)
	return nil
}

func (artC *Prod) DeleteCategories(ctx context.Context, CategorieValue string) error {
	if CategorieValue == "" {
		return fmt.Errorf("CategorieValue is empty")
	}
	query := `DELETE FROM Categories WHERE categorie_value = $1`
	row, err := artC.DB.Exec(query, CategorieValue)
	if err != nil {
		return err
	}
	fmt.Print(row)
	return nil
}

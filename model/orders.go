package model

import (
	"context"
	"fmt"
	"log"

	"github/think.com/dots"
)

// --> CreateSize creates a new Size for a product.
func (artC *Prod) addSizeOrderItem(ctx context.Context, sizes []string, orderitem_id string) error {

	if orderitem_id == "" {
		return nil
	}

	if len(sizes) == 0 {
		return nil

	}

	query := `INSERT INTO sizeorderitem (orderitem_id ,size_value ) VALUES ($1, $2)`

	for _, size := range sizes {

		if size == "" {
			return fmt.Errorf("size is empty")
		}
		log.Print(size)
		_, err := artC.DB.ExecContext(ctx, query, orderitem_id, size)
		if err != nil {
			log.Printf("Filesd to sert: %s for the product_id:%s || %v \n", size, orderitem_id, err)
			return fmt.Errorf("Error adding size : %w", err)
		}
	}

	return nil
}

// --> CreateColor creates a new Color for a product.
func (artC *Prod) addColorOrderItem(ctx context.Context, colors []string, orderitem_id string) error {

	if orderitem_id == "" {
		return nil
	}

	query := `INSERT INTO ColorOrderItem (orderitem_id, color_value) VALUES ($1, $2)`

	for _, color := range colors {

		if color == "" {
			continue
		}
		_, err := artC.DB.ExecContext(ctx, query, orderitem_id, color)
		if err != nil {
			log.Printf("Filesd to sert: %s for the product_id:%s || %v \n", color, orderitem_id, err)
			return fmt.Errorf("Error adding color : %w", err)
		}
	}

	return nil
}

// --> addOrderItem creates a new Color for a product.
func (artC *Prod) addOrderItem(ctx context.Context, OrderItems []dots.OrderItem, OrderID string) error {
	if len(OrderItems) != 0 {
		for _, OrderItem := range OrderItems {
			OrderItemID := artC.generateuild()

			if OrderItem.ProductID == "" {
				return nil
			}

			query := `INSERT INTO OrderItem (orderitem_id , product_id ,order_id, quantity) VALUES ($1, $2 , $3 , $4)`

			_, err := artC.DB.ExecContext(ctx, query, &OrderItemID, OrderItem.ProductID, OrderID, OrderItem.Quantity)
			if err != nil {
				return fmt.Errorf("Error adding addOrderItem : %w ", err)
			}

			if err := artC.addSizeOrderItem(ctx, OrderItem.Sizes, OrderItemID); err != nil {
				return err
			}

			if err := artC.addColorOrderItem(ctx, OrderItem.Colors, OrderItemID); err != nil {
				return err
			}
		}
	}

	return nil
}

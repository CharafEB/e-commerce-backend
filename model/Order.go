package model

import (
	"context"
	"database/sql"
	"fmt"
	"github/think.com/dots"
)

// Getdots.Orders :Get all details for the order
func (artC *Prod) GetOrders(ctx context.Context) ([]dots.Orders, error) {
	// استعلام الطلبات الأساسية
	OrdersQuery := `
		SELECT
			order_id,
			full_name,
			user_state,
			user_city,
			full_address,
			phone_number,
			total_price
		FROM Orders
		ORDER BY order_id
	`

	rows, err := artC.DB.QueryContext(ctx, OrdersQuery)
	if err != nil {
		return nil, fmt.Errorf("Error getting orders: %v", err)
	}
	defer rows.Close()

	var Orders []dots.Orders
	orderMap := make(map[string]*dots.Orders)

	// قراءة الطلبات الأساسية
	for rows.Next() {
		// فحص إلغاء السياق
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var order dots.Orders
		err := rows.Scan(
			&order.OrderID,
			&order.UserFullName,
			&order.UserState,
			&order.UserCity,
			&order.FullAddress,
			&order.PhoneNUmber,
			&order.TotalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("Error reading the order: %v", err)
		}

		order.OrderItem = []dots.OrderItem{}
		Orders = append(Orders, order)
		orderMap[order.OrderID] = &Orders[len(Orders)-1]
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error in reading orders: %v", err)
	}

	// إذا لم توجد طلبات، إرجاع مصفوفة فارغة
	if len(Orders) == 0 {
		return Orders, nil
	}

	// استعلام عناصر الطلبات مع تصفية حسب order_id
	orderItemsQuery := `
		SELECT
			orderitem_id,
			product_id,
			order_id,
			quantity
		FROM orderitem
		ORDER BY order_id, orderitem_id
	`

	args := make([]interface{}, len(Orders))
	for i, order := range Orders {
		args[i] = order.OrderID
	}

	orderItemRows, err := artC.DB.QueryContext(ctx, orderItemsQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("Error getting the OrderItems: %v", err)
	}
	defer orderItemRows.Close()

	orderItemMap := make(map[string]*dots.OrderItem)

	// قراءة عناصر الطلبات
	for orderItemRows.Next() {
		// فحص إلغاء السياق
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var item dots.OrderItem
		err := orderItemRows.Scan(
			&item.OrderItem,
			&item.ProductID,
			&item.OrderID,
			&item.Quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("Error reading order item: %v", err)
		}

		item.Sizes = []string{}
		item.Colors = []string{}

		if order, exists := orderMap[item.OrderID]; exists {
			order.OrderItem = append(order.OrderItem, item)
			orderItemMap[item.OrderItem] = &order.OrderItem[len(order.OrderItem)-1]
		}
	}

	if err = orderItemRows.Err(); err != nil {
		return nil, fmt.Errorf("Error processing order items: %v", err)
	}

	// جلب الأحجام والألوان إذا وجدت عناصر طلبات
	if len(orderItemMap) > 0 {
		orderItemIDs := make([]string, 0, len(orderItemMap))
		for id := range orderItemMap {
			orderItemIDs = append(orderItemIDs, id)
		}

		// جلب الأحجام
		sizesQuery := `
			SELECT
				orderitem_id,
				size_value
			FROM size_order_items
			ORDER BY orderitem_id
		`

		sizeArgs := make([]interface{}, len(orderItemIDs))
		for i, id := range orderItemIDs {
			sizeArgs[i] = id
		}

		sizeRows, err := artC.DB.QueryContext(ctx, sizesQuery, sizeArgs...)
		if err != nil {
			return nil, fmt.Errorf("Error getting the Sizes: %v", err)
		}
		defer sizeRows.Close()

		for sizeRows.Next() {
			// فحص إلغاء السياق
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			var orderItemID, sizeValue string
			err := sizeRows.Scan(&orderItemID, &sizeValue)
			if err != nil {
				return nil, fmt.Errorf("Error reading sizes: %v", err)
			}

			if item, exists := orderItemMap[orderItemID]; exists {
				item.Sizes = append(item.Sizes, sizeValue)
			}
		}

		// جلب الألوان
		colorsQuery := `
			SELECT
				orderitem_id,
				color_value
			FROM color_order_items
			ORDER BY orderitem_id
		`

		colorRows, err := artC.DB.QueryContext(ctx, colorsQuery, sizeArgs...)
		if err != nil {
			return nil, fmt.Errorf("Error getting colors: %v", err)
		}
		defer colorRows.Close()

		for colorRows.Next() {
			// فحص إلغاء السياق
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			var orderItemID, colorValue string
			err := colorRows.Scan(&orderItemID, &colorValue)
			if err != nil {
				return nil, fmt.Errorf("Error reading colors: %v", err)
			}

			if item, exists := orderItemMap[orderItemID]; exists {
				item.Colors = append(item.Colors, colorValue)
			}
		}
	}

	return Orders, nil
}

func (artC *Prod) GetOrderByID(ctx context.Context, orderID string) (*dots.Orders, error) {
	// استعلام الطلب الأساسي
	orderQuery := `
		SELECT
			order_id,
			full_name,
			user_state,
			user_city,
			full_address,
			phone_number,
			total_price
		FROM Orders
		WHERE order_id = ?
	`

	var order dots.Orders
	err := artC.DB.QueryRowContext(ctx, orderQuery, orderID).Scan(
		&order.OrderID,
		&order.UserFullName,
		&order.UserState,
		&order.UserCity,
		&order.FullAddress,
		&order.PhoneNUmber,
		&order.TotalPrice,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Order not found")
		}
		return nil, fmt.Errorf("Error getting order: %v", err)
	}

	order.OrderItem = []dots.OrderItem{}

	// جلب عناصر الطلب
	orderItemsQuery := `
		SELECT
			orderitem_id,
			product_id,
			order_id,
			quantity
		FROM orderitem
		WHERE order_id = ?
		ORDER BY orderitem_id
	`

	orderItemRows, err := artC.DB.QueryContext(ctx, orderItemsQuery, orderID)
	if err != nil {
		return nil, fmt.Errorf("Error getting order items: %v", err)
	}
	defer orderItemRows.Close()

	orderItemMap := make(map[string]*dots.OrderItem)

	for orderItemRows.Next() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var item dots.OrderItem
		err := orderItemRows.Scan(
			&item.OrderItem,
			&item.ProductID,
			&item.OrderID,
			&item.Quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("Error reading order item: %v", err)
		}

		item.Sizes = []string{}
		item.Colors = []string{}

		order.OrderItem = append(order.OrderItem, item)
		orderItemMap[item.OrderItem] = &order.OrderItem[len(order.OrderItem)-1]
	}

	// جلب الأحجام والألوان إذا وجدت عناصر
	if len(orderItemMap) > 0 {
		orderItemIDs := make([]string, 0, len(orderItemMap))
		for id := range orderItemMap {
			orderItemIDs = append(orderItemIDs, id)
		}

		// جلب الأحجام
		sizesQuery := `
			SELECT
				orderitem_id,
				size_value
			FROM size_order_items
		`

		sizeArgs := make([]interface{}, len(orderItemIDs))
		for i, id := range orderItemIDs {
			sizeArgs[i] = id
		}

		sizeRows, err := artC.DB.QueryContext(ctx, sizesQuery, sizeArgs...)
		if err != nil {
			return nil, fmt.Errorf("Error getting sizes: %v", err)
		}
		defer sizeRows.Close()

		for sizeRows.Next() {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			var orderItemID, sizeValue string
			err := sizeRows.Scan(&orderItemID, &sizeValue)
			if err != nil {
				return nil, fmt.Errorf("Error reading size: %v", err)
			}

			if item, exists := orderItemMap[orderItemID]; exists {
				item.Sizes = append(item.Sizes, sizeValue)
			}
		}

		// جلب الألوان
		colorsQuery := `
			SELECT
				orderitem_id,
				color_value
			FROM color_order_items
		`

		colorRows, err := artC.DB.QueryContext(ctx, colorsQuery, sizeArgs...)
		if err != nil {
			return nil, fmt.Errorf("Error getting colors: %v", err)
		}
		defer colorRows.Close()

		for colorRows.Next() {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			var orderItemID, colorValue string
			err := colorRows.Scan(&orderItemID, &colorValue)
			if err != nil {
				return nil, fmt.Errorf("Error reading color: %v", err)
			}

			if item, exists := orderItemMap[orderItemID]; exists {
				item.Colors = append(item.Colors, colorValue)
			}
		}
	}

	return &order, nil
}

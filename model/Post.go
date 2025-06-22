package model

import (
	"context"
	"fmt"
	"github/think.com/dots"
	"github/think.com/migler/index"
	"log"

	"github.com/blevesearch/bleve"
)

// --> Creating the Product
func (artC *Prod) CreateProducts(ctx context.Context, prod *dots.Product, Res bleve.Index) error {
	var val = dots.Products{
		ProductsName:        *&prod.ProductsName,
		ProductsShortDes:    *&prod.ProductsShortDes,
		ProductsDescription: *&prod.ProductsDescription,
		Price:               *&prod.Price,
		ImgURl:              *&prod.ImgURl,
		Categories:          *&prod.Categories,
	}
	val.ProductsID = artC.generateuild()

	query := `INSERT INTO Products (product_id, prod_name, prod_min_dec, prod_doc, prod_price , main_img , categories) VALUES ($1, $2 , $3, $4 , $5 , $6 ,$7)`

	_, err := artC.DB.ExecContext(ctx, query, val.ProductsID, val.ProductsName, val.ProductsShortDes, val.ProductsDescription, val.Price, val.ImgURl, val.Categories)
	if err != nil {
		log.Printf("there is an error adding new Product , ProductsID : %s || Error : %v", val.ProductsID, err)
		return err
	}

	// Indexing the prodect
	if err := index.AssNewIndex(val, Res); err != nil {
		log.Print(err)
	}

	//add colors
	if err := artC.addProductsColor(ctx, prod.Color, val.ProductsID); err != nil {
		return err
	}
	//add imgs
	if err := artC.addProductsImges(ctx, prod.ImgesURL, val.ProductsID); err != nil {
		return err
	}
	//add sizes
	if err := artC.addProductsSize(ctx, prod.Size, val.ProductsID); err != nil {
		return err
	}

	return nil
}

// --> Creating Nrw Order
func (artC *Prod) CreateOrder(ctx context.Context, order *dots.Orders, Res bleve.Index) error {
	var val = dots.Orders{
		UserFullName: *&order.UserFullName,
		UserState:    *&order.UserState,
		UserCity:     *&order.UserCity,
		FullAddress:  *&order.FullAddress,
		TotalPrice:   *&order.TotalPrice,
		PhoneNUmber:  *&order.PhoneNUmber,
	}
	val.OrderID = artC.generateuild()
	query := `INSERT INTO Orders (order_id, full_name, user_state, user_city, full_address , total_price , phone_number) VALUES ($1, $2 , $3, $4 , $5 , $6 , $7)`

	_, err := artC.DB.ExecContext(ctx, query, &val.OrderID, &val.UserFullName, &val.UserState, &val.UserCity, &val.FullAddress, &val.TotalPrice, &val.PhoneNUmber)
	if err != nil {
		log.Printf("there is an error adding new order , ProductsID : %s || Error : %v", val.OrderID, err)
		return err
	}
	fmt.Println(val)
	if err := artC.addOrderItem(ctx, order.OrderItem, val.OrderID); err != nil {
		return err
	}

	return nil
}

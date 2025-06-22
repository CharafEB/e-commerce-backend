package model

import (
	"context"
	"fmt"
	"log"

	"github/think.com/dots"
)

// --> This will handel getting the data to handle it using bleve lib
func (artC *Prod) GetProducts(ctx context.Context) ([]dots.Products, error) {
	rows, err := artC.DB.QueryContext(ctx, "SELECT product_id , prod_name , prod_min_dec  , main_img ,prod_price , categories FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []dots.Products
	for rows.Next() {
		var product dots.Products
		err = rows.Scan(&product.ProductsID, &product.ProductsName, &product.ProductsShortDes, &product.ImgURl, &product.Price, &product.Categories)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// This section will handel getting the content of poroduct from Products , ProductsSize , ProductsColor , ProductsImges
func (artC *Prod) GetProductContent(ctx context.Context, productID string) (*dots.Product, error) {
	//to save data
	var productDetails dots.Product
	log.Println("we are in func")
	productQuery := `
        SELECT *
        FROM products
        WHERE product_id = $1`

	err := artC.DB.QueryRowContext(ctx, productQuery, productID).Scan(
		&productDetails.ProductsID,
		&productDetails.ProductsName,
		&productDetails.ProductsShortDes,
		&productDetails.ProductsDescription,
		&productDetails.Price,
		&productDetails.ImgURl,
		&productDetails.Categories,
	)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}

	sizesQuery := `
        SELECT size_value
        FROM ProductsSize
        WHERE product_id = $1`

	rows, err := artC.DB.QueryContext(ctx, sizesQuery, productID)
	if err != nil {
		return nil, fmt.Errorf("Error in Size: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var size string
		if err := rows.Scan(&size); err != nil {
			log.Printf("Error in Size %v", err)
			continue
		}
		productDetails.Size = append(productDetails.Size, size)
	}

	colorsQuery := `
        SELECT color_value
        FROM ProductsColor
        WHERE product_id = $1`

	rows, err = artC.DB.QueryContext(ctx, colorsQuery, productID)
	if err != nil {
		return nil, fmt.Errorf("Error in Colors: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var color string
		if err := rows.Scan(&color); err != nil {
			log.Printf("Error in colors :%v", err)
			continue
		}
		productDetails.Color = append(productDetails.Color, color)
	}

	imagesQuery := `
       SELECT i.img_value
FROM ProductsImges i
JOIN products p ON i.product_id = p.product_id
WHERE i.img_value != p.main_img AND i.product_id = $1`

	rows, err = artC.DB.QueryContext(ctx, imagesQuery, productID)
	if err != nil {
		return nil, fmt.Errorf("Faild geting img's: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		if err := rows.Scan(&image); err != nil {
			log.Printf("Faild reading img's: %v", err)
			continue
		}
		productDetails.ImgesURL = append(productDetails.ImgesURL, image)
	}
	log.Println(productDetails)

	return &productDetails, nil
}

func (artC *Prod) GetContentSearch(ctx context.Context, val string) ([]dots.Products, error) {
	rows, err := artC.DB.QueryContext(ctx, "SELECT product_id, prod_name, prod_min_dec, prod_doc, prod_price ,main_img , categories FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prods []dots.Products
	for rows.Next() {
		var prod dots.Products
		err = rows.Scan(&prod.ProductsID, &prod.ProductsName, &prod.ProductsShortDes, &prod.ProductsDescription, &prod.Price, &prod.ImgURl, &prod.Categories)
		if err != nil {
			return nil, err
		}
		prods = append(prods, prod)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prods, nil
}

// GetByCategorie retrieves products by category.
func (artC *Prod) GetByCategorie(ctx context.Context, category string) ([]dots.Products, error) {
	query := `
		SELECT product_id, prod_name, prod_min_dec, main_img, categories
		FROM Products
		WHERE categories = $1`

	rows, err := artC.DB.QueryContext(ctx, query, category)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving products by category: %w", err)
	}
	defer rows.Close()

	var products []dots.Products
	for rows.Next() {
		var product dots.Products
		err = rows.Scan(&product.ProductsID, &product.ProductsName, &product.ProductsShortDes, &product.ImgURl, &product.Categories)
		if err != nil {
			return nil, fmt.Errorf("Error scanning product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}

	return products, nil
}

// GetByPriceUnder retrieves products with a price under the specified value.
func (artC *Prod) GetByPriceUnder(ctx context.Context, priceUnder float64) ([]dots.Products, error) {
	query := `
		SELECT product_id, prod_name, prod_min_dec, main_img, categories
		FROM Products
		WHERE prod_price < $1`

	rows, err := artC.DB.QueryContext(ctx, query, priceUnder)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving products under price: %w", err)
	}
	defer rows.Close()

	var products []dots.Products
	for rows.Next() {
		var product dots.Products
		err = rows.Scan(&product.ProductsID, &product.ProductsName, &product.ProductsShortDes, &product.ImgURl, &product.Categories)
		if err != nil {
			return nil, fmt.Errorf("Error scanning product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}

	return products, nil
}

// GetByPriceUper retrieves products with a price over the specified value.
func (artC *Prod) GetByPriceUper(ctx context.Context, priceUper float64) ([]dots.Products, error) {
	query := `
		SELECT product_id, prod_name, prod_min_dec, main_img, categories
		FROM Products
		WHERE prod_price > $1`

	rows, err := artC.DB.QueryContext(ctx, query, priceUper)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving products over price: %w", err)
	}
	defer rows.Close()

	var products []dots.Products
	for rows.Next() {
		var product dots.Products
		err = rows.Scan(&product.ProductsID, &product.ProductsName, &product.ProductsShortDes, &product.ImgURl, &product.Categories)
		if err != nil {
			return nil, fmt.Errorf("Error scanning product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}

	return products, nil
}

// --> GetCategories: will handel getting the data to handle it using bleve lib
func (artC *Prod) GetCategories(ctx context.Context) ([]dots.Categories, error) {
	rows, err := artC.DB.QueryContext(ctx, `SELECT categorie_value FROM Categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []dots.Categories
	for rows.Next() {
		var categorie dots.Categories
		err = rows.Scan(&categorie.CategorieValue)
		if err != nil {
			return nil, err
		}
		fmt.Print(categorie.CategorieValue)
		categories = append(categories, categorie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

// --> GetSize: will handel getting the data to handle it using bleve lib
func (artC *Prod) GetSizes(ctx context.Context) ([]dots.Size, error) {
	rows, err := artC.DB.QueryContext(ctx, `SELECT size_value FROM Size`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Sizes []dots.Size
	for rows.Next() {
		var Size dots.Size
		err = rows.Scan(&Size.SizeValue)
		if err != nil {
			return nil, err
		}
		fmt.Print(Size.SizeValue)
		Sizes = append(Sizes, Size)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return Sizes, nil
}

// --> GetColor: will handel getting the data to handle it using bleve lib
func (artC *Prod) GetColors(ctx context.Context) ([]dots.Color, error) {
	rows, err := artC.DB.QueryContext(ctx, `SELECT color_value FROM color`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colors []dots.Color
	for rows.Next() {
		var color dots.Color
		err = rows.Scan(&color.ColorValue)
		if err != nil {
			return nil, err
		}
		fmt.Print(color.ColorValue)
		colors = append(colors, color)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return colors, nil
}

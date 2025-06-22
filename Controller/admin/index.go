package admin

import (
	middlewares "github/think.com/Middlewares"
	"net/http"
)

type Application struct {
	middlewares.Application
}
type AdminController interface {
	//CreatNewProduct: hander creating new Products on product table
	CreatNewProduct(w http.ResponseWriter, r *http.Request)

	//UpdatProduct: hander creating new Products on product table
	UpdatProduct(w http.ResponseWriter, r *http.Request)

	//ModifierProduct: hander creating new Products on product table
	ModifierProduct(w http.ResponseWriter, r *http.Request)

	//DeleteProduct: hander creating new Products on product table
	DeleteProduct(w http.ResponseWriter, r *http.Request)

	//UploadProductsImgs: hander creating new Products on product table
	UploadProductsImgs(w http.ResponseWriter, r *http.Request)

	//getImageHandler: hander getting the img's
	GetImageHandler(w http.ResponseWriter, r *http.Request)

	//CreatNewCategorie: hander Adding new NewCategorie
	CreatNewCategorie(w http.ResponseWriter, r *http.Request)

	// CreatNewSize: Create a new Size to hte Size table
	CreatNewSize(w http.ResponseWriter, r *http.Request)

	// CreatNewColor: Create a new Color to hte Color table
	CreatNewColor(w http.ResponseWriter, r *http.Request)

	// GetCategories : get the Categories from GetCategories table
	GetCategories(w http.ResponseWriter, r *http.Request)

	// GetSizes : get the Sizes from Size table
	GetSizes(w http.ResponseWriter, r *http.Request)

	// GetColors : get the Sizes from Color table
	GetColors(w http.ResponseWriter, r *http.Request)

	DeleteSize(w http.ResponseWriter, r *http.Request)
	DeleteCategories(w http.ResponseWriter, r *http.Request)
	DeleteColor(w http.ResponseWriter, r *http.Request)

	GetOrdersById(w http.ResponseWriter, r *http.Request)
	GetOrders(w http.ResponseWriter, r *http.Request)
}

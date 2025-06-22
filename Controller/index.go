package controller

import (
	middlewares "github/think.com/Middlewares"
	"net/http"
)

type Application struct {
	middlewares.Application
}

type UserController interface {
	Search(w http.ResponseWriter, r *http.Request)
	CreateNewOrder(w http.ResponseWriter, r *http.Request)
}

type ProductController interface {
	GetProductsByCategory(w http.ResponseWriter, r *http.Request)
	GetProductsByUPrice(w http.ResponseWriter, r *http.Request)
	GetProductsByOPrice(w http.ResponseWriter, r *http.Request)
	GetProductContent(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
}

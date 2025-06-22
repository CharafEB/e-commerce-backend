package router

import (
	"github.com/go-chi/chi/v5"
)

func (user *Control) Product(r chi.Router) chi.Router {

	//This section will handel the search in the program
	r.Route("/product", func(r chi.Router) {
		r.Route("/Categorie", func(r chi.Router) { r.Get("/{category}", user.Controller.GetProductsByCategory) }) // /product/Categorie/{value}

		r.Route("/Uprice", func(r chi.Router) { r.Get("/{price}", user.Controller.GetProductsByUPrice) }) // /product/UPrice/{value}

		r.Route("/Oprice", func(r chi.Router) { r.Get("/{price}", user.Controller.GetProductsByOPrice) }) // /product/OPrice/{value}
		r.Post("/order", user.Controller.CreateNewOrder)
		r.Get("/getallproducts", user.Controller.GetAllProducts) // /product/content
		r.Route("/contentprod", func(r chi.Router) {
			r.Get("/{id}", user.Controller.GetProductContent)// /product/contentprod/{productID}
		})
	})

	return r
}

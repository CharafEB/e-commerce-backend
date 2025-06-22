package router

import (
	"github.com/go-chi/chi/v5"
)

func (user *Control) Admin(r chi.Router) chi.Router {

	r.Route("/Products", func(r chi.Router) {
		r.Post("/", user.AdminController.CreatNewProduct)           // POST /Products
		r.Put("/", user.AdminController.UpdatProduct)               // GET /Products/Updater
		r.Delete("/{id}", user.AdminController.DeleteProduct)       // DELETE /Products/{id}
		r.Post("/imgs", user.AdminController.UploadProductsImgs)    // POST /Products/imgs
		r.Get("/{imageName}", user.AdminController.GetImageHandler) // GET /Products/imgs/{id}
	})

	r.Route("/Categorie", func(r chi.Router) {
		r.Get("/", user.AdminController.GetCategories)
		r.Post("/", user.AdminController.CreatNewCategorie)
		r.Delete("/{q}", user.AdminController.DeleteCategories)
	})

	r.Route("/Size", func(r chi.Router) {
		r.Get("/", user.AdminController.GetSizes)
		r.Post("/", user.AdminController.CreatNewSize)
		r.Delete("/{q}", user.AdminController.DeleteSize)
	})

	r.Route("/Colore", func(r chi.Router) {
		r.Get("/", user.AdminController.GetColors)
		r.Post("/", user.AdminController.CreatNewColor)
		r.Delete("/{q}", user.AdminController.DeleteColor)
	})

	r.Route("/Orders", func(r chi.Router) {
		r.Get("/", user.AdminController.GetOrders)
		r.Get("/{id}", user.AdminController.GetOrdersById)
	})

	return r

}

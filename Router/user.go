package router

import (
	"github.com/go-chi/chi/v5"
)

func (user *Control) User(r chi.Router) chi.Router {

	//This section will handel the search in the program

	r.Route("/Search", func(r chi.Router) {
		r.Get("/{q}", user.Controller.Search) // /Search/{value}
	})
	

	//This ganna set an Order
	 // /setorder

	return r
}

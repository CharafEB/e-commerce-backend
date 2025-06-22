package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (user *Control) Moul() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Mount("/", user.Product(r))
	r.Mount("/Admin", user.Admin(r))
	r.Mount("/user", user.User(r))

	return r

}

func (app *Application) Run(mux *chi.Mux) error {
	var handler http.Handler = mux

	if app.CORSMiddleware != nil {
		handler = app.CORSMiddleware.Handler(mux)
	}

	srv := &http.Server{
		Addr:         app.Address,
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is starting on Port %s", app.Address)

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Server failed to start: %v", err)
		return err
	}
	return nil
}

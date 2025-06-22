package router

import (
	controller "github/think.com/Controller"
	"github/think.com/Controller/admin"
	middlewares "github/think.com/Middlewares"

	"github.com/rs/cors"
)

type Application struct {
	middlewares.Application
	CORSMiddleware *cors.Cors
}

type Controller struct {
	UserControllers    controller.UserController
	ProductControllers controller.ProductController
}

type Admin struct {
	AdminController admin.AdminController
}

type Control struct {
	Controller      *controller.Application
	AdminController *admin.Application
}

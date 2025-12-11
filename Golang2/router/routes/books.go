package routes

import (
	"Golang/controller"
	"Golang/middlewares" 
	"net/http"
)

var booksRoutes = []Route{
	{
		URI:    "/books",
		Method: http.MethodGet,
		Func:   middlewares.Autenticar(controller.HandleSearch), 
	},
}
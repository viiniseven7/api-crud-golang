package routes

import (
	"Golang/controller"
	"net/http"
)
var usersRoutes = []Route{

	{
		URI:    "/login",
		Method: http.MethodPost,
		Func:   controller.Login,
	},
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controller.GetAllUsers,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodGet,
		Func:   controller.GetUser,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
	},
}

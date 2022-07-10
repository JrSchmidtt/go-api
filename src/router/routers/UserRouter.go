package routers

import (
	"api/src/controllers"
	"net/http"
)

var userRouter = []Router{
	{
		Uri:      "/user",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		ReqAuth:  false,
	},
	{
		Uri:      "/user",
		Method:   http.MethodGet,
		Function: controllers.FindAllUsers,
		ReqAuth:  false,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodGet,
		Function: controllers.FindUser,
		ReqAuth:  false,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		ReqAuth:  false,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		ReqAuth:  false,
	},
}

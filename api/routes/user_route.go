package router

import (
	"github.com/gorilla/mux"
	controller "github.com/omjogani/bookmark-manager/controllers"
)

func UserRouter() *mux.Router {
	userRouter := mux.NewRouter()

	userRouter.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	userRouter.HandleFunc("/login", controller.LoginUser).Methods("POST")
	return userRouter
}

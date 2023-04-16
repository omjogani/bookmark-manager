package router

import (
	"github.com/gorilla/mux"
	controller "github.com/omjogani/bookmark-manager/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// userRouter := router.PathPrefix("/api/user").Subrouter()
	// userRouter.StrictSlash(true)

	// UserRouter(userRouter)
	router.HandleFunc("/api/user/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/api/user/logout", controller.LogOut).Methods("POST")
	return router
}

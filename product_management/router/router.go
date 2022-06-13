package router

import (
	"product_management/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// setting up routes

	// public routes login and signup

	router.HandleFunc("/register",controllers.CreateNewUser).Methods("POST") 
	
	router.HandleFunc("/login",controllers.LoginUser).Methods("POST")



	// authorized routes with middleware

	router.HandleFunc("/addProduct",controllers.IsAuthorized(controllers.AddNewProduct)).Methods("POST")

	router.HandleFunc("/getProducts",controllers.IsAuthorized(controllers.GetProducts)).Methods("GET")

	router.HandleFunc("/deleteProducts/{productID}",controllers.IsAuthorized(controllers.DeleteProduct)).Methods("DELETE")
	router.HandleFunc("/updateProduct/{productID}",controllers.IsAuthorized(controllers.UpdateProduct)).Methods("PUT")

	

	return router
}
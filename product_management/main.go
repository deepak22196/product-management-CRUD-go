package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"product_management/database"
	"product_management/router"

	"github.com/gorilla/handlers"
)


func main(){

	database.InitiateDB()
	fmt.Println("server is getting started")
	r:=router.Router()
	
	fmt.Println("router done")
	origins := handlers.AllowedOrigins([]string{"*"})
	headers:=handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization","token"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT","DELETE" ,"OPTIONS"})

	port:=os.Getenv("PORT")
	if port==""{
		port="8000"
	}
	
	log.Fatal(http.ListenAndServe(":"+port,handlers.CORS(origins,headers,methods)(r)))
	
	fmt.Println("dfd")
	fmt.Println("listening at port 8000")

}
//
//


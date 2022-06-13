package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product_management/database"

	"product_management/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// function to add new product to db
func AddNewProduct(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var newProduct models.Product
	err:=json.NewDecoder(r.Body).Decode(&newProduct)
	if err!=nil{
		log.Fatal(err)
	}
	
	newProduct.Creator=r.Header.Get("email")


	inserted,err:=database.ProductCollection.InsertOne(context.Background(),newProduct)
	if err!=nil{
		log.Fatal(err)
	}

	// returning new product id as json response
	fmt.Println("inserted one product in db",inserted.InsertedID)
	json.NewEncoder(w).Encode(newProduct)

}

//
//
//
// function to send all the products for that particular logeed in user
func GetProducts(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","GET")
	var email=r.Header.Get("email")

	// getting email from request header set by middleware, checking products having creator same as logged in user
	cursor,err:=database.ProductCollection.Find(context.Background(),bson.M{"creator":email})
	if err!=nil{
		log.Fatal(err)
	}
	
	var products [] bson.M

	for cursor.Next(context.Background()){
		var product bson.M
		err:=cursor.Decode(&product)
		if err!=nil{
			log.Fatal(err)
		}
		products=append(products,product)
	}
// returning the array as json
	json.NewEncoder(w).Encode(products)

}
//
//
//function to delete the product using product id
func DeleteProduct(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	// getting product id from url params 
	productID:=params["productID"]
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","Delete")
	// converting primitive id type to string
	id,_:=primitive.ObjectIDFromHex(productID)
	filter:=bson.M{"_id":id}
	deleteCount,err:=database.ProductCollection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("number of product deleted:",deleteCount)
	// returning the number of items deleted, which will be one
	json.NewEncoder(w).Encode(productID)
}

//function to delete the product using product id
func UpdateProduct(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	// getting product id from url params 
	productID:=params["productID"]
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	var UpdatedProduct models.Product
	err:=json.NewDecoder(r.Body).Decode(&UpdatedProduct)
	if err != nil {
		log.Fatal(err)
	}
	// converting primitive id type to string
	id,_:=primitive.ObjectIDFromHex(productID)
	// setting the filter and update option for updateOne()
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"brand":UpdatedProduct.Brand,"processor":UpdatedProduct.Processor,"color":UpdatedProduct.Color,"hardisk":UpdatedProduct.Hardisk,"os":UpdatedProduct.Os}}
	updateCount,err:=database.ProductCollection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("number of product deleted:",updateCount)
	json.NewEncoder(w).Encode(productID)
}
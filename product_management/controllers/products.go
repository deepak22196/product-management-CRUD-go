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
	fmt.Println("inserted one product in db",inserted.InsertedID)
	json.NewEncoder(w).Encode(newProduct)

}

//
//
//
func GetProducts(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","GET")
	var email=r.Header.Get("email")
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

	json.NewEncoder(w).Encode(products)

}
//
//
//
func DeleteProduct(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	productID:=params["productID"]
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","Delete")
	id,_:=primitive.ObjectIDFromHex(productID)
	filter:=bson.M{"_id":id}
	deleteCount,err:=database.ProductCollection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("number of product deleted:",deleteCount)
	json.NewEncoder(w).Encode(productID)
}


func UpdateProduct(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	productID:=params["productID"]
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	var UpdatedProduct models.Product
	err:=json.NewDecoder(r.Body).Decode(&UpdatedProduct)
	if err != nil {
		log.Fatal(err)
	}
	id,_:=primitive.ObjectIDFromHex(productID)
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
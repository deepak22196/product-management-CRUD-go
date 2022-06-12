package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Deepak:dreamers111@cluster0.ljtwi.mongodb.net/?retryWrites=true&w=majority"

 const dbName="product_management"
var UserCollection *mongo.Collection
var ProductCollection *mongo.Collection


//  var collection *mongo.Collection
// 
// 
func InitiateDB(){
	
	
clientOption := options.Client().ApplyURI(connectionString)
// ctx:=context.TODO()

client,err:=mongo.Connect(context.TODO(),clientOption)
if err!=nil{
	log.Fatal(err)
}

fmt.Println("database connected successfully")
Product_management_db:=client.Database(dbName)
UserCollection=Product_management_db.Collection("users")
ProductCollection=Product_management_db.Collection("products")
// fmt.Printf("%T%T",UserCollection,ProductCollection)
fmt.Println("connect db completed")

}
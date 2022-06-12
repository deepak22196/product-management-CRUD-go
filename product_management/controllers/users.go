package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product_management/database"
	"time"

	"product_management/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

  var Secretkey string = "secretkeyjwt"
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
//
//
//
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

//take password as input and generate new hash password from it
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)



	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
//
//
//

func CreateNewUser(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var newUser models.User
	err:=json.NewDecoder(r.Body).Decode(&newUser)
	if err!=nil{
		log.Fatal(err)
	}

	//
	var existingUser models.User
	database.UserCollection.FindOne(context.Background(),bson.M{"email":newUser.Email}).Decode(&existingUser)

	if newUser.Email==existingUser.Email{
		w.WriteHeader(http.StatusBadRequest)
		resp:=make(map[string]string)
		resp["message"]="Email Id already registered...please login or use diffrent Id"
		jsonResp,err:=json.Marshal(resp)
		if err!=nil{
			log.Fatal(err)
		}
		w.Write(jsonResp)
		return
		
	}




	//

	newUser.Password, _ = GeneratehashPassword(newUser.Password)

	inserted,err:=database.UserCollection.InsertOne(context.Background(),newUser)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("inserted one user in db",inserted.InsertedID)
	json.NewEncoder(w).Encode(newUser)

}
//
//
//


func LoginUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var authDetails Authentication
	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	database.UserCollection.FindOne(context.Background(),bson.M{"email":authDetails.Email}).Decode(&user)
	if user.Email==""{
		w.WriteHeader(http.StatusBadRequest)
		resp:=make(map[string]string)
		resp["message"]="Email Id not registered...please register first"
		jsonResp,err:=json.Marshal(resp)
		if err!=nil{
			log.Fatal(err)
		}
		w.Write(jsonResp)
		return
		
	}
	check := CheckPasswordHash(authDetails.Password, user.Password)
	if !check{
		w.WriteHeader(http.StatusBadRequest)
		resp:=make(map[string]string)
		resp["message"]="incorrect password"
		jsonResp,err:=json.Marshal(resp)
		if err!=nil{
			log.Fatal(err)
		}
		w.Write(jsonResp)
		return
	}

	validToken, err := GenerateJWT(authDetails.Email)
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token Token
	token.Email = authDetails.Email
	
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
	

}
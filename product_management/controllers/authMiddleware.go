package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// middleware ti check the token, if everything fine then pass on to controllers

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
// checking if there is token header present in request or not
		if r.Header["Token"] == nil {
			var err Error
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(Secretkey)
		// parsing the token 
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Println(err)
			var err Error
			err = SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		// extracting email from token claims and setting it on request header for controllers
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			
				r.Header.Set("email", claims["email"].(string))
				handler.ServeHTTP(w, r)
				return

		
		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}}
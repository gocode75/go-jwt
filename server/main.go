package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("QuickBrownFoxJumpsOverTheLazyDog") // testing
// var mySigningKey = []byte(os.Getenv("MY_JWT_SIGNING_KEY")) // production

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Token Error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequest() {
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("JWT Server")
	handleRequest()
}

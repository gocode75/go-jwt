package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("QuickBrownFoxJumpsOverTheLazyDog") // testing
// var mySigningKey = []byte(os.Getenv("MY_JWT_SIGNING_KEY")) // production

func homepage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Manoj Gupta"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("Error generating token string: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
func main() {
	fmt.Println("JWT Client")

	handleRequest()
}

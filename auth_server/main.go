package main

import (
	"fmt"
	"github.com/hadyelzayady/jwt"
	"net/http"
	"sync"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type datastore struct {
	m map[string]user
	*sync.RWMutex
}

type signupForm struct {
	userName string
	email    string
	password string
}

func login(response http.ResponseWriter, request *http.Request) {
	name := request.Form.Get("email")
	fmt.Fprintf(response, name)
}

type signupRequest struct {
	email    string `json:email`
	password string `json:password`
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func getForm(request *http.Request) signupRequest {

	email := request.FormValue("email")

	password := request.FormValue("password")

	return signupRequest{
		email:    email,
		password: password,
	}

}

//TODO hel
func signup(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()

	token, err := jwt.CreateToken(&jwt.Payload{})

	if err == nil {
		fmt.Fprintf(response, token)
	}
}

func main() {

	http.HandleFunc("/signup", signup)
	// http.HandleFunc("/login", login)

	fmt.Println(http.ListenAndServe(":80", nil))
}

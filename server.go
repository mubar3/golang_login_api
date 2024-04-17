package main

import (
	"fmt"
	"net/http"
	"os"

	Controller "golang_api_login/function"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// routes
	r := mux.NewRouter()
	r.HandleFunc("/", Controller.HomeHandler)
	r.HandleFunc("/hello", Controller.HelloHandler)
	r.HandleFunc("/user/{name}", Controller.UserHandler)
	r.HandleFunc("/get", Controller.Get)
	r.HandleFunc("/post_json", Controller.PostJson)
	r.HandleFunc("/post_form", Controller.PostForm)

	// Menjalankan server dengan menggunakan router Gorilla Mux
	godotenv.Load(".env")
	fmt.Println("Server is listening on port", os.Getenv("PORT_SERVER"))
	http.ListenAndServe(":"+os.Getenv("PORT_SERVER"), r)
}

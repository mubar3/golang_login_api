package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	Controller "golang_api_login/function"
	Model "golang_api_login/model"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Membuka file untuk logging
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	// init db
	Model.Init()

	// routes
	r := mux.NewRouter()
	r.HandleFunc("/", Controller.HomeHandler)
	r.HandleFunc("/hello", Controller.HelloHandler)
	r.HandleFunc("/user/{name}", Controller.UserHandler)
	r.HandleFunc("/get", Controller.Get)
	r.HandleFunc("/post_json", Controller.PostJson)
	r.HandleFunc("/post_form", Controller.PostForm)
	r.HandleFunc("/login_file", Controller.Login_file)
	r.HandleFunc("/login_db", Controller.Login_db)

	// Menjalankan server dengan menggunakan router Gorilla Mux
	godotenv.Load(".env")
	fmt.Println("Server is listening on port", os.Getenv("PORT_SERVER"))
	http.ListenAndServe(":"+os.Getenv("PORT_SERVER"), r)
}

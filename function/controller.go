package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s!", name)
}

func Get(w http.ResponseWriter, r *http.Request) {

	// Create response object
	response := APIResponse{Message: "done"}

	// Convert response object to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func PostJson(w http.ResponseWriter, r *http.Request) {
	// Baca data dari badan permintaan (request body)
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lakukan sesuatu dengan data yang diterima
	// Misalnya, cetak pesan ke konsol
	// println("Received message:", postData.Message)

	// Kirim respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postData)
}

func PostForm(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the value of the "message" field
	message := r.FormValue("message")

	// Create response object
	response := APIResponse{Message: message}

	// Convert response object to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

package Controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Get(w http.ResponseWriter, r *http.Request) {

	// Create response object
	response := APIResponse{Status: true}

	// Convert response object to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	}

	// send respons
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func PostJson(w http.ResponseWriter, r *http.Request) {

	type PostData struct {
		Message string `json:"message"`
	}

	var postData PostData

	json.NewDecoder(r.Body).Decode(&postData)
	responseJSON, _ := json.Marshal(APIResponse{Status: true, Message: postData.Message})

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)

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
	response := APIResponse{Status: true, Message: message}

	// Convert response object to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send respons
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

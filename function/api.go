package Controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

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

func Login(w http.ResponseWriter, r *http.Request) {
	// type output
	w.Header().Set("Content-Type", "application/json")

	// check form
	err := r.ParseForm()
	if err != nil {
		responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: "form eror:" + err.Error()})
		w.Write(responseJSON)
		return
	}

	// read file
	type Data struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	file, err := os.Open("user.json")
	if err != nil {
		responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: "Error load file:" + err.Error()})
		w.Write(responseJSON)
		return
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var data []Data
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: "Error opening file:" + err.Error()})
		w.Write(responseJSON)
		return
	}

	// check user & pass
	status_akun := []any{}
	cek_user := []any{}
	cek_password := []any{}
	for _, item := range data {
		if r.FormValue("user") == item.User && r.FormValue("password") == item.Password {
			status_akun = append(status_akun, true)
		}
		if r.FormValue("user") == item.User {
			cek_user = append(cek_user, true)
		}
		if r.FormValue("password") == item.Password {
			cek_password = append(cek_password, true)
		}
	}

	// output
	if isInSlice(status_akun, true) {
		responseJSON, _ := json.Marshal(APIResponse{Status: true, Message: "login done"})
		w.Write(responseJSON)
	} else {
		if isInSlice(cek_user, true) {
			responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: "wrong password"})
			w.Write(responseJSON)
			return
		}
		if isInSlice(cek_password, true) {
			responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: "wrong user"})
			w.Write(responseJSON)
			return
		}

		// if uknown user & pass
		responseJSON, _ := json.Marshal(APIResponse{Status: false, Message: "login failed, unknown user"})
		w.Write(responseJSON)
	}

}

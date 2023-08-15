package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gksingh1910/go-lang-interface/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("user created with id %v", r.Body)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userModel struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		PhoneNo string `json:"phone_no"`
	}
	b := json.NewDecoder(r.Body)
	b.DisallowUnknownFields()
	b.Decode(&userModel)
	fmt.Printf("Request body is %v", userModel)
	util.RespondWithJson(w, 200, userModel)
}

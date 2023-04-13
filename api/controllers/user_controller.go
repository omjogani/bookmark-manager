package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/omjogani/bookmark-manager/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	error := json.NewDecoder(r.Body).Decode(&user)
	checkNilError(error)

	fmt.Println(user)

	// TODO: Integrate RegisterUser here
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	error := json.NewDecoder(r.Body).Decode(&user)
	checkNilError(error)

	// TODO: Integrate LoginUser here
}

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

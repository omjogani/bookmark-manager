package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	helper "github.com/omjogani/bookmark-manager/helpers"
	model "github.com/omjogani/bookmark-manager/models"
	"github.com/omjogani/bookmark-manager/validations"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	var receivedData model.User
	error := json.NewDecoder(r.Body).Decode(&receivedData)
	checkNilError(error)

	user.Email = receivedData.Email
	user.Password = receivedData.Password
	user.UserName = receivedData.UserName

	fmt.Println(user)
	responseFromValidation := validations.IsUserValid(user)
	if responseFromValidation == "OK" {
		helper.RegisterUserHelper(user)
		json.NewEncoder(w).Encode("User Registered Successfully!")
	} else {
		json.NewEncoder(w).Encode(responseFromValidation)
	}
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

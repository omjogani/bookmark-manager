package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
	var receivedData model.User
	error := json.NewDecoder(r.Body).Decode(&receivedData)
	checkNilError(error)

	user.Email = receivedData.Email
	user.Password = receivedData.Password
	user.UserName = "Demo"

	responseFromValidation := validations.IsUserValid(user)
	if responseFromValidation == "OK" {
		responseData := helper.LoginUserHelper(user)
		expiryFromResponseData, err := strconv.ParseInt(responseData["expiresIn"], 10, 64)
		checkNilError(err)
		expiry := time.Unix(expiryFromResponseData, 0)

		// Set Cookie for Authorization
		cookie := http.Cookie{}
		cookie.Name = "accessToken"
		cookie.Value = responseData["accessToken"]
		cookie.Expires = expiry
		cookie.Secure = true
		cookie.SameSite = http.SameSiteStrictMode
		http.SetCookie(w, &cookie)

		json.NewEncoder(w).Encode(responseFromValidation)
	} else {
		json.NewEncoder(w).Encode(responseFromValidation)
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	accessToken, err := r.Cookie("accesstoken")
	if err != nil {
		log.Fatal("Hello World", err)
	}
	// checkNilError(err)

	// Get UserToken from Cookie
	helper.LogOutUserHelper(accessToken.Value)

	// Delete Cookie
	cookie := http.Cookie{}
	cookie.Name = "accesstoken"
	cookie.Value = ""
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	cookie.MaxAge = -1
	http.SetCookie(w, &cookie)
}

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

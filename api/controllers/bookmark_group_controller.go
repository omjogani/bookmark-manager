package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/omjogani/bookmark-manager/models"
	"github.com/omjogani/bookmark-manager/validations"
)

func AddBookmarkGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencod")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var receivedBookmarkGroup model.BookmarkGroup
	var bookmarkGroup model.BookmarkGroup
	err := json.NewDecoder(r.Body).Decode(&receivedBookmarkGroup)
	checkNilError(err)

	bookmarkGroup.BookmarkGroupName = receivedBookmarkGroup.BookmarkGroupName
	bookmarkGroup.BookmarkGroupDescription = receivedBookmarkGroup.BookmarkGroupDescription
	bookmarkGroup.BookmarkItems = []model.BookmarkItem{}

	responseFromValidation := validations.IsBookmarkGroupValid(bookmarkGroup)
	if responseFromValidation == "OK" {
		// TODO: Call Controller & Return Encoder Response

	} else {
		json.NewEncoder(w).Encode(responseFromValidation)
	}
}

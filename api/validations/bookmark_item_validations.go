package validations

import (
	"regexp"

	model "github.com/omjogani/bookmark-manager/models"
)

func IsBookmarkItemValid(bookmarkItem model.BookmarkItem) string {
	isStringNotEmptyRegEx := regexp.MustCompile(`/^(?!\s*$).+/`)

	if !isStringNotEmptyRegEx.MatchString(bookmarkItem.Title) {
		return "Bookmark Title is required!"
	}

	if !isStringNotEmptyRegEx.MatchString(bookmarkItem.URL) {
		return "URL is required!"
	}
	return "OK"
}

package validations

import (
	"regexp"

	model "github.com/omjogani/bookmark-manager/models"
)

func IsBookmarkGroupValid(bookmarkGroup model.BookmarkGroup) string {
	isStringNotEmptyRegEx := regexp.MustCompile(`/^(?!\s*$).+/`)

	if !isStringNotEmptyRegEx.MatchString(bookmarkGroup.BookmarkGroupName) {
		return "BookmarkGroup Title is required!"
	}
	return "OK"
}

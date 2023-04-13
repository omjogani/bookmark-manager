package model

type BookmarkGroup struct {
	BookmarkGroupId          int            `json:"id"`
	BookmarkGroupName        string         `json:"name"`
	BookmarkGroupDescription string         `json:"description"`
	BookmarkItems            []BookmarkItem `json:"bookmark_items"`
}

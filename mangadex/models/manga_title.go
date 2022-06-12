package models

type MangaOutput struct {
	Result string
	Data   mangaTitleData
}

type mangaTitleData struct {
	Id         int
	Type       int
	Attributes mangaTitleAttributes
}

type mangaTitleAttributes struct {
	Title map[string]string
}

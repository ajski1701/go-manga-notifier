package feed

type FeedOutput struct {
	Results []Results `json:"result"`
	Limit   int
	Offset  int
	Total   int
	Data    []Data
}

type Results struct {
	Result string
}

type Relationships struct {
	Id   string
	Type string
}

type Data struct {
	Id            string
	Type          string
	Attributes    Attributes
	Relationships []Relationships `json:"relationships"`
}

type Attributes struct {
	Title              string
	Volume             string
	Chapter            string
	TranslatedLanguage string `json:"translatedLanguage"`
	Hash               string
	Uploader           string
	ExternalUrl        string
	Version            int
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	PublishAt          string `json:"publishAt"`
}

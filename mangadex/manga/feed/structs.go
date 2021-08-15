package feed

type FeedOutput struct {
	Results []Results `json:"results"`
	Limit   int
	Offset  int
	Total   int
}

type Results struct {
	Result        string
	Data          Data
	Relationships []Relationships `json:"relationships"`
}

type Relationships struct {
	Id   string
	Type string
}

type Data struct {
	Id         string
	Type       string
	Attributes Attributes
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

package structs

type FeedOutput struct {
	Results []feedResults `json:"result"`
	Limit   int
	Offset  int
	Total   int
	Data    []feedData
}

type feedResults struct {
	Result string
}

type feedRelationships struct {
	Id   string
	Type string
}

type feedData struct {
	Id            string
	Type          string
	Attributes    feedAttributes
	Relationships []feedRelationships `json:"relationships"`
}

type feedAttributes struct {
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

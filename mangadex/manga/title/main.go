package title

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetTitle(mangaId string) string {
	url := "https://api.mangadex.org/manga/" + mangaId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result MangaOutput
	json.Unmarshal([]byte(body), &result)
	titleLanguage := ""
	for element := range result.Data.Attributes.Title {
		titleLanguage = element
	}
	return result.Data.Attributes.Title[titleLanguage]
}

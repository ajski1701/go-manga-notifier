package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/ini.v1"
)

func GetAuth(cfg *ini.File) string {
	username := cfg.Section("mangadex").Key("username").String()
	password := cfg.Section("mangadex").Key("password").String()
	values := map[string]string{"username": username, "password": password}

	jsonValue, _ := json.Marshal(values)
	//https://api.mangadex.org/docs.html#operation/post-auth-login
	resp, err := http.Post("https://api.mangadex.org/auth/login", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Authentication Response status:", resp.Status)

	if resp.Status != "200 OK" {
		fmt.Println("Authentication failed. Exiting.")
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var result AuthOutput
	json.Unmarshal([]byte(body), &result)
	sessionToken := result.Token.Session

	return sessionToken
}

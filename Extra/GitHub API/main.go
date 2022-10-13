package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type user struct {
	Login    string
	Name     string
	ReposCnt int `json:"public_repos"`
	Location string
	Bio      string
}

func main() {
	user, err := userInfo("alextldr")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf("%#v\n", user)
}

func userInfo(login string) (*user, error) {
	u := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	user := user{Login: login}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

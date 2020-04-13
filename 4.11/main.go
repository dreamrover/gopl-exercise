package main

import "net/http"

const (
	url   = "https://api.github.com"
	token = "6259026933fcb45e54cd607abb58b4ce6092ea8e"
)

func main() {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+token)
	resp, err := client.Do(req)
	defer resp.Body.Close()
}

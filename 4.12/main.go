package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

const (
	url    = "https://xkcd.com/"
	suffix = "/info.0.json"
	file   = "xkcd.json"
	//vps    = "45.32.54.159"
	//port   = 6379
)

func main() {
	var comic []Comic
	var c Comic
	var keyword string

	info, err := os.Stat(file)
	if err != nil || info.Size() == 0 {
		for i := 1; ; i++ {
			if i == 404 {
				continue
			}
			resp, err := http.Get(url + strconv.Itoa(i) + suffix)
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				break
			}
			err = json.NewDecoder(resp.Body).Decode(&c)
			if err != nil {
				continue
			}
			comic = append(comic, c)
			fmt.Printf("\rScraping data from %s%d", url, i)
		}
		fmt.Printf("\n%d comics saved\n", len(comic))
		data, err := json.MarshalIndent(comic, "", "    ")
		err = ioutil.WriteFile(file, data, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
			return
		}
		if err = json.Unmarshal(data, &comic); err != nil {
			log.Fatal(err)
			return
		}
	}

	for {
		fmt.Print("\n> ")
		fmt.Scanf("%s", &keyword)
		for _, c = range comic {
			if strings.Contains(strings.ToLower(c.Title), strings.ToLower(keyword)) {
				fmt.Printf("%s: %s\n", c.Title, url+strconv.Itoa(c.Num))
				fmt.Println("-------------------------------------------------")
			}
		}
	}
}

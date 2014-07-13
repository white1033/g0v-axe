package main

import (
	"encoding/json"
	"fmt"
	"github.com/Puerkitobio/goquery"
	"net/http"
	"net/http/cookiejar"
)

type 里長 struct {
	Town    string `json:"town"`
	Village string `json:"village"`
	Name    string `json:"name"`
}

func main() {
	var result []里長
	var url = "http://axe-level-1.herokuapp.com/lv3"
	options := cookiejar.Options{PublicSuffixList: nil}
	jar, _ := cookiejar.New(&options)
	client := http.Client{Jar: jar}
	for i := 1; i < 77; i += 1 {
		resp, _ := client.Get(url)
		url = "http://axe-level-1.herokuapp.com/lv3?page=next"
		doc, _ := goquery.NewDocumentFromResponse(resp)
		doc.Find("table.table tr").Each(func(i int, s *goquery.Selection) {
			if i > 0 {
				info := s.Children().Map(func(i int, s *goquery.Selection) string { return s.Text() })
				result = append(result, 里長{info[0], info[1], info[2]})
			}
		})
	}
	js, _ := json.Marshal(result)
	fmt.Printf("%s", js)
}

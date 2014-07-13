package main

import (
	"encoding/json"
	"fmt"
	"github.com/Puerkitobio/goquery"
)

type 里長 struct {
	Town    string `json:"town"`
	Village string `json:"village"`
	Name    string `json:"name"`
}

func main() {
	var result []里長
	var page = 1
	var url = "http://axe-level-1.herokuapp.com/lv2?page=%d"

	for page < 13 {
		doc, _ := goquery.NewDocument(fmt.Sprintf(url, page))
		doc.Find("table.table tr").Each(func(i int, s *goquery.Selection) {
			if i > 0 {
				info := s.Children().Map(func(i int, s *goquery.Selection) string { return s.Text() })
				result = append(result, 里長{info[0], info[1], info[2]})
			}
		})
		page += 1
	}

	js, _ := json.Marshal(result)
	fmt.Printf("%s", js)
}

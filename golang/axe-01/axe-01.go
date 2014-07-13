package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type Person struct {
	Name   string `json:"name"`
	Grades Grade  `json:"grades"`
}

type Grade struct {
	Chinese          int `json:"國語"`
	Math             int `json:"數學"`
	Science          int `json:"自然"`
	Sociology        int `json:"社會"`
	HeathEducational int `json:"健康教育"`
}

func main() {
	var result []Person
	var grades = make([]int, 5)
	var name string

	doc, _ := goquery.NewDocument("http://axe-level-1.herokuapp.com/")

	doc.Find("table.table tr").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			s.Children().Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					name = s.Text()
				} else {
					grades[i-1], _ = strconv.Atoi(s.Text())
				}
			})
			grade := &Grade{grades[0], grades[1], grades[2], grades[3], grades[4]}
			person := Person{name, *grade}
			result = append(result, person)
		}
	})
	js, _ := json.Marshal(result)
	fmt.Printf("%s", js)
}

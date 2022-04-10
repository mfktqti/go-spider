package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func goQuery() {
	url := "https://gorm.io/zh_CN/docs"

	d, _ := goquery.NewDocument(url)
	d.Find(".sidebar-link").Each(
		func(i int, s *goquery.Selection) {
			// s2 := s.Text()
			// fmt.Printf("s2: %v\n", s2)
			href, _ := s.Attr("href")
			fmt.Printf("href: %v\n", href)
		})

	goQuery1()
}

func goQuery1() {
	html := `
	<body>
	<div>DIV1</div>
	<div>DIV2</div>
	<span>SPAN</span>
	</body>`
	d, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	d.Find("span").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("s.Text(): %v\n", s.Text())
	})
}

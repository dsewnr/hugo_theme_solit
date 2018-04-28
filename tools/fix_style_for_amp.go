package main

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	arg := os.Args[1]
	f, e := os.Open(arg)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	style := doc.Find("head style[amp-custom]").First()
	doc.Find("body style").Each(func(i int, s *goquery.Selection) {
		style.AppendHtml(s.Text())
		s.Remove()
	})

	fo, err := os.Create(arg + ".new")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	html, _ := doc.Html()

	fo.WriteString(html)

	os.Rename(arg+".new", arg)
}

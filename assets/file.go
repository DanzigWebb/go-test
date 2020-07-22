package assets

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
)

func GetHtml(src string) *os.File {
	file, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func GetQueryDoc(src string) *goquery.Document {
	file := GetHtml(src)
	doc, err := goquery.NewDocumentFromReader(file)

	if err != nil {
		log.Fatal(err)
	}

	return doc
}
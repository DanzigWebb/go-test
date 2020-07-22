package assets

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// GetQueryDoc return queryHTML
func GetQueryDoc(src string) *goquery.Document {
	file := getHTML(src)
	doc, err := goquery.NewDocumentFromReader(file)

	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func getHTML(src string) *os.File {
	file, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

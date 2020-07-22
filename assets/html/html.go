package html

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

// HTML ...
type HTML struct {
	query      *goquery.Document
	bundleName string
	selector   string
}

// New HTML config
func New(query *goquery.Document, bundleName, Selector string) *HTML {
	return &HTML{
		query:      query,
		bundleName: bundleName,
		selector:   Selector,
	}
}

// Query return goQuery html
func (s *HTML) Query() *goquery.Document {
	return s.query
}

// GetOutputHTML return html file afte modify
func (s *HTML) GetOutputHTML(cutScripts string) string {
	body := s.query.Find("body")
	body.SetHtml(s.getBodyAndScripts(cutScripts) + "\n")
	outputHTML, err := s.query.Html()

	if err != nil {
		log.Fatal(err)
	}

	return outputHTML
}

// GetJsScript return js script for render html
func (s *HTML) GetJsScript() string {
	return fmt.Sprintf("document.querySelector('#%s').innerHTML = file", s.selector)
}

func (s *HTML) getBodyAndScripts(cutScripts string) string {
	content := s.getBody(s.bundleName) + cutScripts
	return content
}

func (s *HTML) getBody(bundleName string) string {
	return fmt.Sprintf("<div id='%s'></div>\n<script src='%s'></script>\n", s.selector, bundleName)
}

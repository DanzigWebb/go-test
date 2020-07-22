package html

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

// HTML ...
type HTML struct {
	bundleName string
	selector   string
}

// New HTML config
func New(bundleName, Selector string) *HTML {
	return &HTML{
		bundleName: bundleName,
		selector:   Selector,
	}
}

// GetOutputHTML return html file afte modify
func (s *HTML) GetOutputHTML(query *goquery.Document, cutScripts string) string {
	body := query.Find("body")
	body.SetHtml(s.getBodyAndScripts(cutScripts) + "\n")
	outputHTML, err := query.Html()

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

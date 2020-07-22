package html

import (
	"fmt"
	"log"
	"spa/assets/vars"

	"github.com/PuerkitoBio/goquery"
)

var bundleName, selector = vars.BundleName, vars.Selector

func GetOutputHtml(query *goquery.Document, bodyScripts string) string {
	body := query.Find("body")
	body.SetHtml(GetBodyAndScripts(bundleName, bodyScripts))
	outputHtml, err := query.Html()

	if err != nil {
		log.Fatal(err)
	}

	return outputHtml
}

func GetJsScript() string {
	return fmt.Sprintf("document.querySelector('#%s').innerHTML = file", selector)
}

func GetBodyAndScripts(bundleName, cutScripts string) string {
	content := getBody(bundleName) + cutScripts
	return content
}

func getBody(bundleName string) string {
	return fmt.Sprintf("<div id='%s'></div> \n <script src='%s'></script>", selector, bundleName)
}

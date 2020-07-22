package bundle

import (
	"fmt"
	"spa/assets"
	"spa/assets/html"

	"github.com/PuerkitoBio/goquery"
)

// CreateBundle ...
func CreateBundle(doc *goquery.Document, HTML *html.HTML) string {

	body := doc.Find("body")
	bundleHTML, _ := body.Html()

	jsContent := fmt.Sprintf("var file = `%s`; \n %s \n", bundleHTML, HTML.GetJsScript())

	return jsContent
}

// GenerateBundleName ...
func GenerateBundleName(length int) string {
	return fmt.Sprintf("%s.bundle.js", assets.CreateRandomStr(length))
}

package bundle

import (
	"fmt"
	"spa/assets"
	"spa/assets/html"

	"github.com/PuerkitoBio/goquery"
)

func CreateBundle(doc *goquery.Document) string {

	body := doc.Find("body")
	bundleHtml, _ := body.Html()

	jsContent := fmt.Sprintf("var file = `%s`; \n %s", bundleHtml, html.GetJsScript())

	return jsContent
}

func GenerateBundleName(length int) string {
	return fmt.Sprintf("%s.bundle.js", assets.CreateRandomStr(length))
}

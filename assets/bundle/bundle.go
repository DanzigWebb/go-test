package bundle

import (
	"fmt"
	"spa/assets/html"
)

// CreateBundle and return it as string
func CreateBundle(HTML *html.HTML) string {

	body := HTML.Query().Find("body")
	bundleHTML, _ := body.Html()

	jsContent := fmt.Sprintf("var file = `%s`; \n %s \n", bundleHTML, HTML.GetJsScript())

	return jsContent
}

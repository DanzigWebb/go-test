package scripts

import (
	"fmt"
	"spa/assets/html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Cut scripts from body and return them
func Cut(HTML *html.HTML) string {
	var scripts = make([]string, 0)

	HTML.Query().Find("body script").Each(func(index int, script *goquery.Selection) {
		src, hasSrc := script.Attr("src")
		content := script.Text()

		scriptAssets := src
		if !hasSrc {
			scriptAssets = content
		}

		scripts = append(scripts, createScript(hasSrc, scriptAssets))

		script.Remove()
	})

	return strings.Join(scripts[:], "\n")
}

func createScript(hasSrc bool, scriptAsset string) string {
	if hasSrc {
		return fmt.Sprintf("<script src='%s'></script>", scriptAsset)
	} else {
		return fmt.Sprintf("<script>%s</script>", scriptAsset)
	}
}

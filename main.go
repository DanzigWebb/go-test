package main

import (
	"io/ioutil"
	"spa/assets"
	"spa/assets/bundle"
	"spa/assets/html"

	"spa/assets/scripts"
)

func main() {

	var (
		bundleName = assets.GenerateBundleName(12)
		selector   = assets.GenerateSelectorName(10)
	)

	query := assets.GetQueryDoc("ru/index.html")
	HTML := html.New(query, bundleName, selector)

	var (
		bodyScripts = scripts.Cut(HTML)
		bundle      = bundle.CreateBundle(HTML)
		outputHTML  = HTML.GetOutputHTML(bodyScripts)
		newHTML     = "ru/copy.html"
		newBundle   = "ru/" + bundleName
	)

	ioutil.WriteFile(newHTML, []byte(outputHTML), 0600)
	ioutil.WriteFile(newBundle, []byte(bundle), 0600)
}

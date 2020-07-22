package main

import (
	"io/ioutil"
	"spa/assets"
	"spa/assets/bundle"
	"spa/assets/html"

	"spa/assets/scripts"
	"spa/assets/vars"
)

func main() {

	vars.BundleName = bundle.GenerateBundleName(12)
	query := assets.GetQueryDoc("ru/index.html")

	HTML := html.New(query, vars.BundleName, vars.Selector)

	bodyScripts := scripts.Cut(HTML)
	bundle := bundle.CreateBundle(HTML)
	outputHTML := HTML.GetOutputHTML(bodyScripts)

	newHTML, newBundle := "ru/copy.html", "ru/"+vars.BundleName

	ioutil.WriteFile(newHTML, []byte(outputHTML), 0600)
	ioutil.WriteFile(newBundle, []byte(bundle), 0600)
}

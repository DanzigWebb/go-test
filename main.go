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

	HTML := html.New(vars.BundleName, vars.Selector)
	query := assets.GetQueryDoc("mock/index.html")

	bodyScripts := scripts.Cut(query)
	bundle := bundle.CreateBundle(query, HTML)
	outputHTML := HTML.GetOutputHTML(query, bodyScripts)

	newHTML, newBundle := "mock/copy.html", "mock/"+vars.BundleName

	ioutil.WriteFile(newHTML, []byte(outputHTML), 0600)
	ioutil.WriteFile(newBundle, []byte(bundle), 0600)
}

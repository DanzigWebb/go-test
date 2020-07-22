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

	query := assets.GetQueryDoc("mock/index.html")

	bodyScripts := scripts.Cut(query)
	bundle := bundle.CreateBundle(query)
	outputHtml := html.GetOutputHtml(query, bodyScripts)

	newHtml, newBundle := "mock/copy.html", "mock/"+vars.BundleName

	ioutil.WriteFile(newHtml, []byte(outputHtml), 0600)
	ioutil.WriteFile(newBundle, []byte(bundle), 0600)
}

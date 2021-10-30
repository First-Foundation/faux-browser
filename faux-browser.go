package main

import (
	faux_browser "github.com/CyberSecurityN00b/faux-browser/pkg/faux-browser"
)

func main() {
	//faux_browser.NewProfile("testing").StartBrowsing()
	faux_browser.NewProfile("").StartBrowsing()
}

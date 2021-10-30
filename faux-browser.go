package main

import (
	faux_browser "github.com/CyberSecurityN00b/faux-browser/pkg/faux-browser"
)

func main() {
	// USER TODO: Modify/add to the below based on your needs. Setting up multiple browsing profiles
	//	will simulate multiple users (although from the host perspective, they will be using the same browser).
	//	Use "testing" as the profile if you are testing things out, as most profiles won't work after hours.

	//faux_browser.NewProfile("testing").StartBrowsing()
	faux_browser.NewProfile("").StartBrowsing()
}

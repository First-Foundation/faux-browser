package faux_browser

import (
	"os"
	"path/filepath"
	"runtime"
)

type Browser struct {
	NavigatableDomains []string
	UserAgent          string
}

func NewBrowser() (b *Browser) {
	b = &Browser{
		UserAgent: "faux-browser v1.0",
	}

	// Change the user agent based on OS and filename
	//	(e.g., different user agent if the filename is "chrome.exe", "firefox.exe")
	browsername := filepath.Base(os.Args[0])
	switch runtime.GOOS {
	case "darwin":
		// TODO: Verify these are the correct names
		switch browsername {
		case "Google Chrome.app":
			b.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
		case "Safari.app":
			b.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15"
		}
	case "linux":
		switch browsername {
		case "chromium":
			b.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
		case "firefox":
			b.UserAgent = "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:24.0) Gecko/20100101 Firefox/24.0"
		}
	case "windows":
		switch browsername {
		case "chrome.exe":
			b.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
		case "edge.exe":
			b.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763"
		case "firefox.exe":
			b.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0"
		case "iexplore.exe":
			b.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko"
		}
	}

	b.NavigatableDomains = []string{
		// Google Sites
		"google.com",
		"www.google.com",

		// Reddit Sites
		"reddit.com",
		"www.reddit.com",

		// Wikipedia Sites
		"wikipedia.org",
		"en.wikipedia.org"}

	return
}

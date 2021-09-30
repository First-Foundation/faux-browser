package faux_browser

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/net/html"
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

func (b *Browser) ConductSearch(s Search, p *Profile) (urls []string, sleeptime int64) {
	// URL holder
	var u string

	// Pick a random query
	if len(s.Queries) <= 0 {
		return
	}
	i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(s.Queries))))
	query := url.QueryEscape(s.Queries[i.Int64()])

	// How many links are we following?
	GetMinMax(s.MinResultsToClick, s.MaxResultsToClick)

	switch s.Engine {
	case SearchEngine_Bing:
		u = fmt.Sprintf("https://www.bing.com/search?q=%s", query)
	case SearchEngine_Google:
		u = fmt.Sprintf("https://google.com/search?q=%s", query)
	case SearchEngine_Yahoo:
		u = fmt.Sprintf("https://search.yahoo.com/search?p=%s", query)
	}

	urls, sleeptime = b.VisitSite(Site{u, SITE_SEARCHENGINE, []string{}, 0, 0}, p)

	return
}

func (b *Browser) VisitSite(s Site, p *Profile) (urls []string, sleeptime int64) {
	url := s.SeedURL

	// Make the request!
	body := GetRequest(url, false)

	// Make the request look real and get those links!
	SimulateBrowserVisit(url, body)

	// TODO: Get links/results

	// TODO: Filter links/results against allow list

	// TODO: Add links/results to urls

	// Before we close out, does it need to be opened in a real browser?
	switch runtime.GOOS {
	case "darwin":
		exec.Command("open", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("cmd", "/C", "start", url).Start()
	}

	return
}

func GetRequest(url string, ignorecontent bool) (body []byte) {
	resp, err := http.Get(url)
	if err != nil {
		// Bail!
		return
	}
	defer resp.Body.Close()
	if !ignorecontent {
		body, _ = io.ReadAll(resp.Body)
	}
	return
}

// Check for web files (.js, .css, images) and request them. Also
// checks for <a href="xxx"> style links and returns them.
func SimulateBrowserVisit(u string, b []byte) (words int64, links []string) {
	tokenizer := html.NewTokenizer(bytes.NewReader((b)))
	simulatedownload := func(attr string) {
		_, err := url.ParseRequestURI(attr)
		if err == nil {
			GetRequest(attr, true)
		} else {
			_, err := url.ParseRequestURI(u + "/" + attr)
			if err == nil {
				GetRequest(u+"/"+attr, true)
			}
		}
	}
tokenloop:
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "a":
				// Add the href attribute to our list of links
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			case "img":
				// Get that image!
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						simulatedownload(attr.Val)
					}
				}
			case "link":
				// Poss stylesheets, get them!
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						simulatedownload(attr.Val)
					}
				}
			case "script":
				// Poss javascript, get the files!
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						simulatedownload(attr.Val)
					}
				}
			}
		case html.TextToken:
			words = words + int64(len(strings.Fields(string(tokenizer.Text()))))
			//DEBUG
			println(tokenizer.Text())
		case html.ErrorToken:
			// Bail if EOF
			if tokenizer.Err() == io.EOF {
				break tokenloop
			}
		}
	}
	return
}

func GetMinMax(min int64, max int64) (result int64) {
	resultbig, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	result = resultbig.Int64()
	if err == nil {
		result = result + min
	} else {
		result = 0
	}
	return
}

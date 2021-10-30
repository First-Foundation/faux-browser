package faux_browser

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	mathrand "math/rand"
	"net/http"
	"net/url"
	urlpkg "net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

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

	// List of sites that can be visited by default, useful for search engine queries
	b.NavigatableDomains = []string{
		// Adobe Sites
		"adobe.com",
		"get.adobe.com",
		"support.apple.com",
		"www.adobe.com",

		// Amazon Sites
		"amazon.com",
		"www.amazon.com",

		// Amazon AWS Sites
		"amazonaws.com",
		"aws.amazon.com",
		"docs.aws.amazon.com",

		// Apache Sites
		"apache.org",
		"community.apache.org",
		"httpd.apache.org",

		// Apple Sites
		"apple.com",
		"developer.apple.com",
		"download.apple.com",
		"guide.apple.com",
		"help.apple.com",
		"info.apple.com",
		"ipod.apple.com",
		"itunes.apple.com",
		"jobs.apple.com",
		"macosx.apple.com",
		"myinfo.apple.com",
		"quicktime.apple.com",
		"retail.apple.com",
		"sales.apple.com",
		"store.apple.com",
		"support.apple.com",
		"train.apple.com",
		"training.apple.com",
		"www.apple.com",

		// Archive.org Sites
		"archive.org",
		"blog.archive.org",
		"help.archive.org",

		// BBC Sites
		"bbc.com",
		"www.bbc.com",

		// BestBuy Sites
		"bestbuy.com",
		"www.bestbuy.com",

		// Bit.ly Sites
		// USER TODO: This is a url shortener, make sure you want this!
		"bit.ly",

		// Blogspot Sites
		"blogspot.com",
		"www.blogspot.com",

		// Booking.com Sites
		"booking.com",
		"www.booking.com",

		// Britannica Sites
		"britannica.com",
		"www.britannica.com",

		// Cambridge Sites
		"cambridge.com",
		"www.cambridge.com",

		// CDC Sites
		"cdc.gov",
		"www.cdc.gov",
		"wwwnc.cdc.gov",

		// CNN Sites
		"cnn.com",
		"www.cnn.com",

		// Craigslist Sites
		"craigslist.org",
		"www.craigslist.org",

		// Dictionary.com Sites
		"dictionary.com",
		"www.dictionary.com",

		// Dominos Sites
		"dominos.com",
		"www.dominos.com",

		// Dropbox Sites
		"dropbox.com",
		"www.dropbox.com",

		// Ebay Sites
		"ebay.com",
		"www.ebay.com",

		// ESPN Sites
		"espn.com",
		"www.espn.com",

		// Facebook Sites
		"facebook.com",
		"www.facebook.com",

		// Flickr Sites
		"flickr.com",
		"www.flickr.com",

		// Forbes Sites
		"forbes.com",
		"www.forbes.com",

		// Fox News Sites
		"foxnews.com",
		"www.foxnews.com",

		// Github Sites
		"github.com",
		"docs.github.com",
		"help.github.com",
		"pages.github.com",
		"www.github.com",

		"github.io",

		// Godaddy Sites
		"godaddy.com",
		"parked-content.godaddy.com",
		"www.godaddy.com",

		// Google Sites
		"google.com",
		"adwords.google.com",
		"ads.google.com",
		"answers.google.com",
		"ap.google.com",
		"artsandculture.google.com",
		"blogsearch.google.com",
		"books.google.com",
		"clients.google.com",
		"clients1.google.com",
		"checkout.google.com",
		"code.google.com",
		"desktop.google.com",
		"developers.google.com",
		"dl.google.com",
		"docs.google.com",
		"drive.google.com",
		"earth.google.com",
		"feedproxy.google.com",
		"finance.google.com",
		"fusion.google.com",
		"gmail.google.com",
		"groups.google.com",
		"images.google.com",
		"mail.google.com",
		"maps.google.com",
		"news.google.com",
		"pack.google.com",
		"partnerpage.google.com",
		"picasa.google.com",
		"picasaweb.google.com",
		"play.google.com",
		"plus.google.com",
		"podcasts.google.com",
		"policies.google.com",
		"scholar.google.com",
		"services.google.com",
		"sites.google.com",
		"sketchup.google.com",
		"spreadsheets.google.com",
		"stadia.google.com",
		"suggestqueries.google.com",
		"support.google.com",
		"talkgadget.google.com",
		"toolbar.google.com",
		"translate.google.com",
		"www.google.com",
		"video.google.com",
		"video-stats.video.google.com",

		"goo.gl",

		"googleblog.com",
		"developers.googleblog.com",

		"googletagmanager.com",
		"www.googletagmanager.com",

		// Gravatar Sites
		"gravatar.com",
		"www.gravatar.com",

		// Healthline Sites
		"healthline.com",
		"www.healthline.com",

		// HomeDepot Sites
		"homedepot.com",
		"www.homedepot.com",

		// HuffPost Sites
		"huffpost.com",
		"www.huffpost.com",

		// IMDB Site
		"imdb.com",
		"www.imdb.com",

		// Indeed.com Sites
		"indeed.com",
		"www.indeed.com",

		// Instagram Sites
		"instagram.com",
		"www.instagram.com",

		// Investopedia Sites
		"investopedia.com",
		"www.investopedia.com",

		// LinkedIn Sites
		"linkedin.com",
		"www.linkedin.com",

		// Livescore Sites
		"livescore.com",
		"www.livescore.com",

		// Macromedia Sites
		"macromedia.com",
		"www.macromedia.com",

		// Mayo Clinic Sites
		"mayoclinic.org",
		"www.mayoclinic.org",

		// Medium.com Sites
		"medium.com",
		"blog.medium.com",
		"www.medium.com",

		"medium.statuspage.io",

		// Medium.com - Populare Blogs
		"doctorow.medium.com",
		"luke.medium.com",
		"williamfleitch.medium.com",

		// Merriam-Webster Sites
		"merriam-webster.com",
		"www.merriam-webster.com",

		// Microsoft Sites
		"microsoft.com",
		"www.microsoft.com",

		// MSNBC Sites
		"msnbc.com",
		"www.msnbc.com",

		// Mozilla Sites
		"mozilla.com",
		"support.mozilla.com",
		"www.mozilla.com",

		// NBC News Sites
		"nbcnews.com",
		"www.nbcnews.com",

		// Netflix Sites
		"netflix.com",
		"www.netflix.com",

		// New York Times Sites
		"nytimes.com",
		"www.nytimes.com",

		// Nordstrom Sites
		"nordstrom.com",
		"www.nordstrom.com",

		// NPR Sites
		"npr.org",
		"www.npr.org",

		// PHP.net Sites
		"php.net",
		"www.php.net",

		// Pinterest Sites
		"pinterest.com",
		"www.pinterest.com",

		// Playstation Sites
		"playstation.com",
		"partners.playstation.net",
		"status.playstation.com",
		"www.playstation.com",

		// Reddit Sites
		"reddit.com",
		"reddit.statuspage.io",
		"www.reddit.com",

		// Roblox Sites
		"roblox.com",
		"corp.roblox.com",
		"www.roblox.com",

		// RottenTomatoes Sites
		"rottentomatoes.com",
		"www.rottentomatoes.com",

		// Soundcloud Sites
		"soundcloud.com",
		"community.soundcloud.com",

		// SourceForge Sites
		"sourceforge.net",

		// Speedtest.net Sites
		"speedtest.net",
		"www.speedtest.net",

		// Spotify Sites
		"spotify.com",
		"www.spotify.com",

		// TheFreeDictionary Sites
		"thefreedictionary.com",
		"www.thefreedictionary.com",

		// TheGuardian Sites
		"theguardian.com",
		"www.theguardian.com",

		// Thesaurus.com Sites
		"thesaurus.com",
		"www.thesaurus.com",

		// TimeAndDate Sites
		"timeanddate.com",
		"www.timeanddate.com",

		// TripAdvisor Sites
		"tripadvisor.com",
		"www.tripadvisor.com",

		// Tumblr Sites
		"tumblr.com",
		"www.tumblr.com",

		// Twitch Sites
		"twitch.tv",
		"dev.twitch.tv",
		"www.twitch.tv",

		"twitchadvertising.tv",

		// Twitter Sites
		"twitter.com",
		"www.twitter.com",

		// Vimeo Sites
		"vimeo.com",
		"player.vimeo.com",
		"www.vimeo.com",

		// W3 Sites
		"w3.org",
		"www.w3.org",

		// Walmart Sites
		"walmart.com",
		"www.walmart.com",

		// Weather.com Sites
		"weather.com",

		// WebMD Sites
		"webmd.com",
		"www.webmd.com",

		// WhatsApp Sites
		"whatsapp.com",
		"www.whatsapp.com",

		// Wikipedia Sites
		"wikipedia.org",
		"en.wikipedia.org",

		// Wordpress Sites
		"wordpress.com",
		"www.wordpress.com",

		// Yahoo Sites
		"yahoo.com",
		"news.yahoo.com",
		"www.yahoo.com",

		// Yelp Sites
		"yelp.com",
		"www.yelp.com",

		// YouTube Sites
		"youtube.com",
		"www.youtube.com",
		"youtu.be"}

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
	case SearchEngine_Indeed:
		u = fmt.Sprintf("https://www.indeed.com/jobs?q=%s", query)
	case SearchEngine_Yahoo:
		u = fmt.Sprintf("https://search.yahoo.com/search?p=%s", query)
	default:
		return
	}

	urls_tmp, sleeptime := b.VisitSite(Site{u, SITE_SEARCHENGINE, []string{}, 0, 0}, p)

	var search_results []string

	switch s.Engine {
	case SearchEngine_Bing:
		// Bing just gives the raw url in links, so filter on what doesn't have bing.
		for _, url := range urls_tmp {
			if !strings.Contains(url, "bing.com") {
				search_results = append(search_results, url)
			}
		}
	case SearchEngine_Google:
		// Filter on "/url?q=" where it does *not* have accounts.google.com in it
		for _, url := range urls_tmp {
			if strings.Contains(url, "google.com/url?q=") && !strings.Contains(url, "accounts.google.com") {
				search_results = append(search_results, url)
			}
		}
	case SearchEngine_Indeed:
		// Filter only on those urls with "viewjob"
		for _, url := range urls_tmp {
			if strings.Contains(url, "indeed.com/viewjob") {
				search_results = append(search_results, url)
			}
		}
	case SearchEngine_Yahoo:
		// Filter on "r.search.yahoo.com", but filter out those redirecting to yahoo
		for _, url := range urls_tmp {
			if strings.Contains(url, "r.search.yahoo.com") && !strings.Contains(url, "yahoo.com%%2f") {
				search_results = append(search_results, url)
			}
		}
	default:
		return
	}

	urls = GetRandomSubSetOfUrls(search_results, s.MinResultsToClick, s.MaxResultsToClick)

	return
}

// VisitSite simulates visiting the site specified, based on the profile provided.
// It returns urls (a list of urls to visit further) and sleeptime (time to wait in seconds).
func (b *Browser) VisitSite(s Site, p *Profile) (urls []string, sleeptime int64) {
	println("Visiting " + s.SeedURL)

	url := s.SeedURL
	var urls_tmp []string

	// Make the request!
	url, body := GetRequest(url, false)

	// Make the request look real and get those links!
	words, links := SimulateBrowserVisit(url, body)

	// Handle the links that were returned
	parsedurl, err := urlpkg.Parse(url)
	if err != nil {
		return
	}
	domain := parsedurl.Hostname()

	for _, l := range links {
		// Modify links that assume the URI (i.g., are 'index.html' instead of 'www.example.com/index.html')
		parsedurl_l, err := urlpkg.ParseRequestURI(l)

		if err != nil {
			break
		}
		if parsedurl_l.Hostname() == "" {
			l = parsedurl.Scheme + "://" + strings.ReplaceAll(domain+"/"+l, "//", "/")
			parsedurl_l, err = urlpkg.Parse(l)
		}

		// Filter against allow list
		if err == nil {
			match_l := false
			domain_l := parsedurl_l.Hostname()

			// Check if it is in the smaller site list of domains we can visit
			for _, d := range s.NavigatableDomains {
				if d == domain_l {
					match_l = true
					break
				}
			}

			// If no match, then check against the larger list
			if !match_l {
				for _, d := range b.NavigatableDomains {
					if d == domain_l {
						match_l = true
						break
					}
				}
			}

			if match_l {
				// At this point, the url is in the allow list
				urls_tmp = append(urls_tmp, l)
			} else {
				println("No allow-list match for: " + l)
			}
		}
	}

	// Before we close out, does it need to be opened in a real browser?
	if s.Options&SITE_USEREALBROWSER > 0 {
		switch runtime.GOOS {
		case "darwin":
			exec.Command("open", url).Start()
		case "linux":
			exec.Command("xdg-open", url).Start()
		case "windows":
			exec.Command("cmd", "/C", "start", url).Start()
		}
	}

	// If this is a SearchEngine search, return everything, else return a random subset.
	if s.Options&SITE_SEARCHENGINE > 0 {
		urls = urls_tmp
	} else {
		urls = GetRandomSubSetOfUrls(urls_tmp, 1, 2)
	}

	// How long to sleep?
	sleeptime = int64(float32(words) * (float32(p.PercentOfSiteRead) / 100.0) / float32(p.GetWPM()*60))

	return
}

// Perform the get request, returning the url (in case of redirects) and data
func GetRequest(url string, ignorecontent bool) (newurl string, body []byte) {
	resp, err := http.Get(url)
	if err != nil {
		// Bail!
		return
	}
	newurl = resp.Request.URL.String()

	defer resp.Body.Close()
	if ignorecontent {
		io.ReadAll(resp.Body)
	} else {
		body, _ = io.ReadAll(resp.Body)
	}
	return
}

// Check for web files (.js, .css, images) and request them. Also
// checks for <a href="xxx"> style links and returns them.
func SimulateBrowserVisit(u string, b []byte) (words int64, links []string) {
	parsedurl, err := urlpkg.Parse(u)
	if err != nil {
		return
	}
	domain := parsedurl.Hostname()
	tokenizer := html.NewTokenizer(bytes.NewReader((b)))

	simulatedownload := func(attr string) {
		_, err := urlpkg.ParseRequestURI(attr)
		if err == nil {
			GetRequest(attr, true)
		} else {
			l := parsedurl.Scheme + "://" + strings.ReplaceAll(domain+"/"+attr, "//", "/")
			_, err := urlpkg.ParseRequestURI(l)
			if err == nil {
				GetRequest(l, true)
			}
		}
	}

	// Use below to track if the last token was one we should ignore
	lasttoken := ""

tokenloop:
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()

		switch tokenType {
		case html.StartTagToken, html.SelfClosingTagToken:
			lasttoken = token.Data
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
				var val string
				var isstylesheet bool
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						val = attr.Val
					} else if attr.Key == "rel" {
						if attr.Val == "stylesheet" {
							isstylesheet = true
						}
					}
				}
				if isstylesheet {
					simulatedownload(val)
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
			if lasttoken != "script" && lasttoken != "noscript" && lasttoken != "style" {
				words = words + int64(len(strings.Fields(token.Data)))
			}
		case html.ErrorToken:
			// Bail if EOF
			if tokenizer.Err() == io.EOF {
				break tokenloop
			}
		}
	}
	println(words)
	return
}

func GetMinMax(min int64, max int64) (result int64) {
	if max < min {
		tmp := max
		max = min
		min = tmp
	}
	if min <= 0 {
		return 0
	}
	resultbig, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	result = resultbig.Int64()
	if err == nil {
		result = result + min
	} else {
		result = 0
	}
	return
}

func GetRandomSubSetOfUrls(urls []string, min int64, max int64) (results []string) {
	if len(urls) == 0 {
		return
	}
	mathrand.Seed(time.Now().Unix())
	for i := 0; i < int(GetMinMax(min, max)); i++ {
		results = append(results, urls[mathrand.Intn(len(urls))])
	}
	return
}

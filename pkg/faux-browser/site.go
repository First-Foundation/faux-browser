package faux_browser

type Site struct {
	SeedURL            string
	Options            SiteOption
	NavigatableDomains []string // In addition to those specified in browser.go
	MinSecondsOnSite   int64    // If 0, Profile's "estimated WPM" result is used as the min
	MaxSecondsOnSite   int64    // If 0, Profile's "estimated WPM" is used as the max
}

type SiteOption byte

const (
	// Don't click any links. Ignored if SITE_CANCLICKANYLINK is specified.
	SITE_DONTCLICKLINKS SiteOption = 0b00000000

	// SITE_CANCLICKLINK will "click" links on the site, provided they match
	// Site.NavigatableDomains
	SITE_CANCLICKLINK SiteOption = 0b00000001

	// SITE_CANCLICKANYLINK will "click" links on the site, provided they match
	// Site.NavigatableDomains -OR- Browser.NavigatableDomains
	SITE_CANCLICKANYLINK SiteOption = 0b00000010

	// SITE_USEREALYBROWSER will attempt to open the site in a real browser.
	// This is not guaranteed to actually happen. Use sparingly, as the browser
	// will not be closed or otherwise interacted with. Useful if you want to
	// actually run javascript, etc. If not set, will "psudeo" open by GET-ing
	// .js, .css, .jpeg, .etc files.
	SITE_USEREALBROWSER SiteOption = 0b00000100

	// SITE_SEARCHENGINE identifies the site as being a search engine. Used to
	// modify the logic to pick the links to visit.
	SITE_SEARCHENGINE SiteOption = 0b00001000
)

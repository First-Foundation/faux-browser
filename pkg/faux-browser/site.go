package faux_browser

type Site struct {
	SeedURL            string
	CanNavigateLinks   bool
	CanNavigateAnyLink bool     // Ignores NavigatableDomains below, but not in browser.go
	NavigatableDomains []string // In addition to those specified in browser.go
	MinSecondsOnSite   int      // If 0, Profile's "estimated WPM" result is used as the min
	MaxSecondsOnSite   int      // If 0, Profile's "estimated WPM" is used as the max
}

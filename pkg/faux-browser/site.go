package faux_browser

type Site struct {
	SeedURL            string
	CanNavigateLinks   bool
	CanNavigateAnyLink bool     // Ignores NavigatableDomains
	NavigatableDomains []string // In addition to those specified in browser.go
}

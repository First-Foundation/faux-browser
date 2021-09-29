package faux_browser

// Generate a profile for someone who likes to research
func (p *Profile) GenerateProfile_Researcher() {
	p.PercentOfSiteRead = 100
	p.WPM = 250
	p.WPM_Jitter = 25

	p.Sites = append(p.Sites, Site{"https://en.wikipedia.org", true, false, []string{"en.wikipedia.org"}, 0, 0})
}

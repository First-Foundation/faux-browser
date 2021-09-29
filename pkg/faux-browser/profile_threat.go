package faux_browser

// Generate a profile of sites that an "unintentional insider threat" would visit
func (p *Profile) GenerateProfile_Threat() {
	p.PercentOfSiteRead = 66
	p.WPM = 150
	p.WPM_Jitter = 25

	// USER TODO: Modify the below to match your needs!
	p.Sites = append(p.Sites, Site{"https://zh.wikipedia.org/", true, false, []string{"zh.wikipedia.org"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://ru.wikipedia.org", true, false, []string{"ru.wikipedia.org"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", false, false, []string{}, 30, 90})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", true, false, []string{"youtube.com", "www.youtube.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"http://www.reddit.com/", true, false, []string{"reddit.com", "www.reddit.com"}, 0, 0})
}

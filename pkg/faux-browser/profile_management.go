package faux_browser

// Generate a profile of sites that a CEO or department lead would visit
func (p *Profile) GenerateProfile_Management() {
	p.PercentOfSiteRead = 66
	p.WPM = 300
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://www.foreignaffairs.com", true, false, []string{"foreignaffairs.com", "www.foreignaffairs.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.ted.com", true, false, []string{"ted.com", "www.ted.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com", true, false, []string{"linkedin.com", "www.linkedin.com"}, 60, 300})
	p.Sites = append(p.Sites, Site{"https://www.ceo.com", true, false, []string{"ceo.com", "www.ceo.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"http://wsj.com/", true, false, []string{"wsj.com", "www.wsj.com"}, 0, 0})
}

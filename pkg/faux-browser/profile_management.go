package faux_browser

// Generate a profile of sites that a CEO or department lead would visit
func (p *Profile) GenerateProfile_Management() {
	p.PercentOfSiteRead = 66
	p.WPM = 300
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://www.foreignaffairs.com", SITE_CANCLICKLINK, []string{"foreignaffairs.com", "www.foreignaffairs.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.ted.com", SITE_CANCLICKLINK, []string{"ted.com", "www.ted.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com", SITE_CANCLICKLINK, []string{"linkedin.com", "www.linkedin.com"}, 60, 300})
	p.Sites = append(p.Sites, Site{"https://www.ceo.com", SITE_CANCLICKLINK, []string{"ceo.com", "www.ceo.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"http://wsj.com/", SITE_CANCLICKLINK, []string{"wsj.com", "www.wsj.com"}, 0, 0})
}

package faux_browser

// Generate a profile of sites that someone in I.T. would visit
func (p *Profile) GenerateProfile_IT() {
	p.PercentOfSiteRead = 50
	p.WPM = 200
	p.WPM_Jitter = 30

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/InformationTechnology/", SITE_CANCLICKLINK, []string{"reddit.com", "www.reddit.com"}, 0, 0})
}

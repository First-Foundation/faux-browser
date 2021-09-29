package faux_browser

// Generate a profile of sites that a "driven" employee would visit
func (p *Profile) GenerateProfile_Employee_Driven() {
	p.PercentOfSiteRead = 80
	p.WPM = 300
	p.WPM_Jitter = 50

	// USER TODO: Add company related sites!
	p.Sites = append(p.Sites, Site{"https://www.themuse.com/advice/47-habits-of-highly-successful-employees", SITE_DONTCLICKLINKS, []string{}, 0, 0})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", SITE_CANCLICKLINK, []string{"youtube.com", "www.youtube.com"}, 60, 180})
}

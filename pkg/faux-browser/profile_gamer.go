package faux_browser

// Generate a profile of sites that a gamer would visit
func (p *Profile) GenerateProfile_Gamer() {
	p.PercentOfSiteRead = 80
	p.WPM = 200
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://store.steampowered.com", SITE_CANCLICKLINK, []string{"steampowered.com", "store.steampowered.com"}, 60, 180})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", SITE_CANCLICKLINK, []string{"youtube.com", "www.youtube.com"}, 300, 600})
}

package faux_browser

// Add common sites that everyone will be visiting
func (p *Profile) GenerateProfile_Common() {
	// DON'T MESS WITH ANYTHING OTHER THAN ADDING TO SITES AND SEARCHES

	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", SITE_CANCLICKLINK, []string{"youtube.com", "www.youtube.com"}, 60, 300})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com/", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"linkedin.com", "www.linkedin.com"}, 300, 0})
	p.Sites = append(p.Sites, Site{"https://www.facebook.com/", SITE_CANCLICKLINK, []string{"facebook.com", "www.facebook.com"}, 60, 180})
}

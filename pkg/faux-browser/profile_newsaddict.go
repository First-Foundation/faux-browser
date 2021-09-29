package faux_browser

// Generate a profile of someone who browsers news websites
func (p *Profile) GenerateProfile_NewsAddict() {
	p.PercentOfSiteRead = 75
	p.WPM = 200
	p.WPM_Jitter = 20

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/news", true, true, []string{}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.nbcnews.com", true, false, []string{"nbcnews.com", "www.nbcnews.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", true, false, []string{"buzzfeed.com", "www.buzzfeed.com"}, 0, 0})
}

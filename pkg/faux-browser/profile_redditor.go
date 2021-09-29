package faux_browser

// Generate a profile of someone who browses reddit at work all day
func (p *Profile) GenerateProfile_Redditor() {
	p.PercentOfSiteRead = 100
	p.WPM = 150
	p.WPM_Jitter = 50

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/AskReddit", true, false, []string{"reddit.com", "www.reddit.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/pics", true, false, []string{"reddit.com", "www.reddit.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", true, false, []string{"buzzfeed.com", "www.buzzfeed.com"}, 0, 0})
}

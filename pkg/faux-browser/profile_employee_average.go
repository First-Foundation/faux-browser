package faux_browser

// Generate a profile of sites that an average employee would visit
func (p *Profile) GenerateProfile_Employee_Average() {
	p.PercentOfSiteRead = 80
	p.WPM = 200
	p.WPM_Jitter = 10

	// USER TODO: Add company related sites!
	p.Sites = append(p.Sites, Site{"https://www.facebook.com/", true, false, []string{"facebook.com", "www.facebook.com"}, 60, 180})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com/", true, false, []string{"linkedin.com", "www.linkedin.com"}, 60, 180})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", true, false, []string{"youtube.com", "www.youtube.com"}, 60, 300})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", true, false, []string{"buzzfeed.com", "www.buzzfeed.com"}, 30, 0})

	p.Searches = append(p.Searches, Search{SearchEngine_Bing, []string{"sports game last week", "ball game last night", "who won the game", "game scores"}, 2, 4})
}

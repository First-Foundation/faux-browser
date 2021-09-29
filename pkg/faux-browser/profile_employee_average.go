package faux_browser

// Generate a profile of sites that an average employee would visit
func (p *Profile) GenerateProfile_Employee_Average() {
	p.PercentOfSiteRead = 80
	p.WPM = 200
	p.WPM_Jitter = 10

	// USER TODO: Add company related sites!
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", SITE_CANCLICKLINK, []string{"buzzfeed.com", "www.buzzfeed.com"}, 30, 0})

	p.Searches = append(p.Searches, Search{SearchEngine_Bing, []string{"sports game last week", "ball game last night", "who won the game", "game scores"}, 2, 4})
}

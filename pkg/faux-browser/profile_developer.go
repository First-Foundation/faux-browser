package faux_browser

// Generate a profile of sites that a Developer might visit
func (p *Profile) GenerateProfile_Developer() {
	p.PercentOfSiteRead = 100
	p.WPM = 200
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://docs.microsoft.com/en-us/", SITE_CANCLICKLINK, []string{"docs.microsoft.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://stackoverflow.com/questions/tagged/c%2B%2B", SITE_CANCLICKLINK, []string{"stackoverflow.com"}, 0, 0})

	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"golang code won't compile", "go.mod import package dissappears on save"}, 1, 3})
}

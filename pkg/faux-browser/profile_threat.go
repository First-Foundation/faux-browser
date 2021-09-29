package faux_browser

// Generate a profile of sites that an "unintentional insider threat" would visit
func (p *Profile) GenerateProfile_Threat() {
	p.PercentOfSiteRead = 66
	p.WPM = 150
	p.WPM_Jitter = 25

	// USER TODO: Modify the below to match your needs!
	p.Sites = append(p.Sites, Site{"https://zh.wikipedia.org/", SITE_CANCLICKLINK, []string{"zh.wikipedia.org"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://ru.wikipedia.org", SITE_CANCLICKLINK, []string{"ru.wikipedia.org"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", SITE_DONTCLICKLINKS, []string{}, 30, 90})
}

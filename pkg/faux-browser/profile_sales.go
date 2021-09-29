package faux_browser

// Generate a profile of sites a sales person would visit
func (p *Profile) GenerateProfile_Sales() {
	p.PercentOfSiteRead = 75
	p.WPM = 225
	p.WPM_Jitter = 25

	p.Sites = append(p.Sites, Site{"https://www.salesforce.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"salesforce.com", "www.salesforce.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.neilpatel.com", SITE_CANCLICKLINK, []string{"neilpatel.com", "www.neilpatel.com"}, 0, 0})
}

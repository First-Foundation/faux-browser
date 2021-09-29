package faux_browser

// Generate a profile of sites that someone in Cybersecurity would visit
func (p *Profile) GenerateProfile_Cybersecurity() {
	p.PercentOfSiteRead = 50
	p.WPM = 250
	p.WPM_Jitter = 50

	p.Sites = append(p.Sites, Site{"https://book.hacktricks.xyz", true, false, []string{"book.hacktricks.xyz"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.virustotal.com", true, false, []string{"virustotal.com", "www.virustotal.com"}, 0, 300})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", false, false, []string{}, 300, 600})

	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"dirtycow exploit", "dirty cow exploit", "proof of concept exploit for dirty cow"}, 1, 5})
}

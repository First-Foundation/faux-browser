package faux_browser

import "time"

// Generate a profile of sites that someone in Cybersecurity would visit
func (p *Profile) GenerateProfile_Cybersecurity() {
	p.ScheduleFunc = ProfileScheduleFunc_Cybersecurity
	p.PercentOfSiteRead = 50
	p.WPM = 250
	p.WPM_Jitter = 50

	p.Sites = append(p.Sites, Site{"https://book.hacktricks.xyz", SITE_CANCLICKLINK, []string{"book.hacktricks.xyz"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.virustotal.com", SITE_CANCLICKLINK, []string{"virustotal.com", "www.virustotal.com"}, 0, 300})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", SITE_DONTCLICKLINKS, []string{}, 300, 600})

	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"dirtycow exploit", "dirty cow exploit", "proof of concept exploit for dirty cow"}, 1, 5})
}

func ProfileScheduleFunc_Cybersecurity(t time.Time) (atdesk bool) {
	// Am I on lunch?
	if t.After(String2Time24("11:00")) && t.Before(String2Time24("12:00")) {
		return false
	}

	// Is this a day of the week I work?
	switch t.Weekday() {
	case time.Monday:
		return true
	case time.Tuesday:
		return true
	case time.Wednesday:
		return true
	case time.Thursday:
		return true
	case time.Friday:
		return true
	case time.Saturday:
		return false
	case time.Sunday:
		return false
	}

	return
}

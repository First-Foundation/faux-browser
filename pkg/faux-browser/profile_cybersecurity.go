package faux_browser

import "time"

// Generate a profile of sites that someone in Cybersecurity would visit
func (p *Profile) GenerateProfile_Cybersecurity() {
	p.ScheduleFunc = ProfileScheduleFunc_Cybersecurity
	p.PercentOfSiteRead = 50
	p.WPM = 250
	p.WPM_Jitter = 50

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"https://book.hacktricks.xyz", SITE_CANCLICKLINK, []string{"book.hacktricks.xyz"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.virustotal.com", SITE_CANCLICKLINK, []string{"virustotal.com", "www.virustotal.com"}, 0, 300})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", SITE_DONTCLICKLINKS, []string{}, 300, 600})
	p.Sites = append(p.Sites, Site{"https://www.exploit-db.com", SITE_CANCLICKLINK, []string{"exploit-db.com", "www.exploit-db.com"}, 30, 90})
	p.Sites = append(p.Sites, Site{"https://nakedsecurity.sophos.com", SITE_CANCLICKANYLINK, []string{"sophos.com", "nakedsecurity.sophos.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.darkreading.com/", SITE_CANCLICKANYLINK, []string{"darkreading.com", "www.darkreading.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.cvedetails.com", SITE_CANCLICKANYLINK, []string{"cvedetails.com", "www.cvedetails.com"}, 30, 0})
	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/netsec", SITE_CANCLICKANYLINK, []string{"reddit.com", "www.reddit.com"}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////

	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"dirtycow exploit", "dirty cow exploit", "proof of concept exploit for dirty cow"}, 1, 5})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"cybersecurity news", "cyber sec news", "infosec news", "hacker news", "hackers in the news"}, 1, 2})

	// List of ransomware from wikipedia
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"reveton", "cryptolocker", "cryptolocker.f", "torrentlocker", "cryptowall", "fusob", "wannacry", "petya", "bad rabbit", "samsam", "darkside", "syskey"}, 1, 2})

	// List of C2 frameworks
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"asyncrat", "caldera", "powershell empire", "covenant", "simple tactical agent relay", "poshc2", "quasar rat", "merlin c2", "metasploit", "meterpreter", "cobalt strike", "voodoo c2"}, 1, 2})
}

func ProfileScheduleFunc_Cybersecurity(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("09:00")) || t.After(String2Time24("17:00")) {
		return false
	}

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

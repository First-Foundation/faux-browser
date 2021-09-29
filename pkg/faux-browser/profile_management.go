package faux_browser

import "time"

// Generate a profile of sites that a CEO or department lead would visit
func (p *Profile) GenerateProfile_Management() {
	p.ScheduleFunc = ProfileScheduleFunc_Management
	p.PercentOfSiteRead = 66
	p.WPM = 300
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://www.foreignaffairs.com", SITE_CANCLICKLINK, []string{"foreignaffairs.com", "www.foreignaffairs.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.ted.com", SITE_CANCLICKLINK, []string{"ted.com", "www.ted.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com", SITE_CANCLICKLINK, []string{"linkedin.com", "www.linkedin.com"}, 60, 300})
	p.Sites = append(p.Sites, Site{"https://www.ceo.com", SITE_CANCLICKLINK, []string{"ceo.com", "www.ceo.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"http://wsj.com/", SITE_CANCLICKLINK, []string{"wsj.com", "www.wsj.com"}, 0, 0})
}

func ProfileScheduleFunc_Management(t time.Time) (atdesk bool) {
	// Am I on lunch?
	if t.After(String2Time24("10:30")) && t.Before(String2Time24("12:30")) {
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

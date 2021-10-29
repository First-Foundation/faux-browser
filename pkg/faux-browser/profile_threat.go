package faux_browser

import "time"

// Generate a profile of sites that an "unintentional insider threat" would visit
func (p *Profile) GenerateProfile_Threat() {
	p.ScheduleFunc = ProfileScheduleFunc_Threat
	p.PercentOfSiteRead = 66
	p.WPM = 150
	p.WPM_Jitter = 25

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	// USER TODO: Modify the below to match your needs!
	p.Sites = append(p.Sites, Site{"https://zh.wikipedia.org/", SITE_CANCLICKLINK, []string{"zh.wikipedia.org"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://ru.wikipedia.org", SITE_CANCLICKLINK, []string{"ru.wikipedia.org"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", SITE_DONTCLICKLINKS, []string{}, 30, 90})
	p.Sites = append(p.Sites, Site{"https://www.exploit-db.com", SITE_CANCLICKLINK, []string{"exploit-db.com", "www.exploit-db.com"}, 30, 90})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////
}

func ProfileScheduleFunc_Threat(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("08:00")) || t.After(String2Time24("16:00")) {
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

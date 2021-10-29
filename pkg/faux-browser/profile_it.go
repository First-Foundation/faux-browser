package faux_browser

import "time"

// Generate a profile of sites that someone in I.T. would visit
func (p *Profile) GenerateProfile_IT() {
	p.ScheduleFunc = ProfileScheduleFunc_IT
	p.PercentOfSiteRead = 50
	p.WPM = 200
	p.WPM_Jitter = 30

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/InformationTechnology/", SITE_CANCLICKLINK, []string{"reddit.com", "www.reddit.com"}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////
}

func ProfileScheduleFunc_IT(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("06:00")) || t.After(String2Time24("18:00")) {
		return false
	}

	// Am I on lunch?
	if t.After(String2Time24("11:30")) && t.Before(String2Time24("12:30")) {
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

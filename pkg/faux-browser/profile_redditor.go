package faux_browser

import "time"

// Generate a profile of someone who browses reddit at work all day
func (p *Profile) GenerateProfile_Redditor() {
	p.ScheduleFunc = ProfileScheduleFunc_Redditor
	p.PercentOfSiteRead = 100
	p.WPM = 150
	p.WPM_Jitter = 50

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/AskReddit", SITE_CANCLICKANYLINK, []string{"reddit.com", "www.reddit.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/pics", SITE_CANCLICKANYLINK, []string{"reddit.com", "www.reddit.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", SITE_CANCLICKANYLINK, []string{"buzzfeed.com", "www.buzzfeed.com"}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////
}

func ProfileScheduleFunc_Redditor(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("08:00")) || t.After(String2Time24("16:30")) {
		return false
	}

	// Am I on lunch?
	if t.After(String2Time24("11:05")) && t.Before(String2Time24("12:05")) {
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

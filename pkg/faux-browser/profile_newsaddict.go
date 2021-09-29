package faux_browser

import "time"

// Generate a profile of someone who browsers news websites
func (p *Profile) GenerateProfile_NewsAddict() {
	p.ScheduleFunc = ProfileScheduleFunc_NewsAddict
	p.PercentOfSiteRead = 75
	p.WPM = 200
	p.WPM_Jitter = 20

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/news", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.nbcnews.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"nbcnews.com", "www.nbcnews.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"buzzfeed.com", "www.buzzfeed.com"}, 0, 0})
}

func ProfileScheduleFunc_NewsAddict(t time.Time) (atdesk bool) {
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

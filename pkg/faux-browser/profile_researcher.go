package faux_browser

import "time"

// Generate a profile for someone who likes to research
func (p *Profile) GenerateProfile_Researcher() {
	p.ScheduleFunc = ProfileScheduleFunc_Researcher
	p.PercentOfSiteRead = 100
	p.WPM = 250
	p.WPM_Jitter = 25

	p.Sites = append(p.Sites, Site{"https://en.wikipedia.org", SITE_CANCLICKLINK, []string{"en.wikipedia.org"}, 0, 0})
}

func ProfileScheduleFunc_Researcher(t time.Time) (atdesk bool) {
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

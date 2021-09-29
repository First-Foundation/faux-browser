package faux_browser

import "time"

// Generate a profile of sites that a gamer would visit
func (p *Profile) GenerateProfile_Gamer() {
	p.ScheduleFunc = ProfileScheduleFunc_Gamer
	p.PercentOfSiteRead = 80
	p.WPM = 200
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://store.steampowered.com", SITE_CANCLICKLINK, []string{"steampowered.com", "store.steampowered.com"}, 60, 180})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", SITE_CANCLICKLINK, []string{"youtube.com", "www.youtube.com"}, 300, 600})
}

func ProfileScheduleFunc_Gamer(t time.Time) (atdesk bool) {
	// Am I on lunch?
	if t.After(String2Time24("11:15")) && t.Before(String2Time24("12:20")) {
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

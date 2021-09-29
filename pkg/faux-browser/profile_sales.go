package faux_browser

import "time"

// Generate a profile of sites a sales person would visit
func (p *Profile) GenerateProfile_Sales() {
	p.ScheduleFunc = ProfileScheduleFunc_Sales
	p.PercentOfSiteRead = 75
	p.WPM = 225
	p.WPM_Jitter = 25

	p.Sites = append(p.Sites, Site{"https://www.salesforce.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"salesforce.com", "www.salesforce.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.neilpatel.com", SITE_CANCLICKLINK, []string{"neilpatel.com", "www.neilpatel.com"}, 0, 0})
}

func ProfileScheduleFunc_Sales(t time.Time) (atdesk bool) {
	// Am I on lunch?
	if t.After(String2Time24("10:00")) && t.Before(String2Time24("11:00")) {
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

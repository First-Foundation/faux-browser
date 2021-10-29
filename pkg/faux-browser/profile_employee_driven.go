package faux_browser

import "time"

// Generate a profile of sites that a "driven" employee would visit
func (p *Profile) GenerateProfile_Employee_Driven() {
	p.ScheduleFunc = ProfileScheduleFunc_EmployeeDriven
	p.PercentOfSiteRead = 80
	p.WPM = 300
	p.WPM_Jitter = 50

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	// USER TODO: Add company related sites!
	p.Sites = append(p.Sites, Site{"https://www.themuse.com/advice/47-habits-of-highly-successful-employees", SITE_DONTCLICKLINKS, []string{}, 0, 0})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", SITE_CANCLICKLINK, []string{"youtube.com", "www.youtube.com"}, 60, 180})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////
}

func ProfileScheduleFunc_EmployeeDriven(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("07:45")) || t.After(String2Time24("17:00")) {
		return false
	}

	// Am I on lunch?
	if t.After(String2Time24("11:10")) && t.Before(String2Time24("12:10")) {
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

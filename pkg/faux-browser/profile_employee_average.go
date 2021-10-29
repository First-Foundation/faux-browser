package faux_browser

import "time"

// Generate a profile of sites that an average employee would visit
func (p *Profile) GenerateProfile_Employee_Average() {
	p.ScheduleFunc = ProfileScheduleFunc_EmployeeAverage
	p.PercentOfSiteRead = 80
	p.WPM = 200
	p.WPM_Jitter = 10

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	// USER TODO: Add company related sites!
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", SITE_CANCLICKLINK, []string{"buzzfeed.com", "www.buzzfeed.com"}, 30, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////

	p.Searches = append(p.Searches, Search{SearchEngine_Bing, []string{"sports game last week", "ball game last night", "who won the game", "game scores"}, 2, 4})
}

func ProfileScheduleFunc_EmployeeAverage(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("08:15")) || t.After(String2Time24("15:45")) {
		return false
	}

	// Am I on lunch?
	if t.After(String2Time24("10:50")) && t.Before(String2Time24("12:10")) {
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

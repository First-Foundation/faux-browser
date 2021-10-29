package faux_browser

import "time"

// Generate a profile of sites that a Developer might visit
func (p *Profile) GenerateProfile_Developer() {
	p.ScheduleFunc = ProfileScheduleFunc_Developer
	p.PercentOfSiteRead = 100
	p.WPM = 200
	p.WPM_Jitter = 10

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"https://docs.microsoft.com/en-us/", SITE_CANCLICKLINK, []string{"docs.microsoft.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://stackoverflow.com/questions/tagged/c%2B%2B", SITE_CANCLICKLINK, []string{"stackoverflow.com"}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////

	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"golang code won't compile", "go.mod import package dissappears on save"}, 1, 3})
}

func ProfileScheduleFunc_Developer(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("08:00")) || t.After(String2Time24("18:00")) {
		return false
	}

	// Am I on lunch?
	if t.After(String2Time24("11:10")) && t.Before(String2Time24("11:50")) {
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

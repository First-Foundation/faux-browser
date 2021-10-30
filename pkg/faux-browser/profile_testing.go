package faux_browser

import "time"

// IMPORTANT PRO-TIP: DON'T NAME THIS FILE 'profile_test.go', AS GO WILL JUST IGNORE IT FOR BUILDING PURPOSES

// Use this profile to conduct testing!
func (p *Profile) GenerateProfile_Testing() {
	p.ScheduleFunc = ProfileScheduleFunc_Testing
	p.PercentOfSiteRead = 50
	p.WPM = 200
	p.WPM_Jitter = 30

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/InformationTechnology/", SITE_CANCLICKLINK, []string{}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////
}

func ProfileScheduleFunc_Testing(t time.Time) (atdesk bool) {
	// Test profile is always ready
	return true
}

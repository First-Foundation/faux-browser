package faux_browser

import "time"

// Generate a profile of someone who browsers news websites
func (p *Profile) GenerateProfile_NewsAddict() {
	p.ScheduleFunc = ProfileScheduleFunc_NewsAddict
	p.PercentOfSiteRead = 75
	p.WPM = 200
	p.WPM_Jitter = 20

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/news", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.nbcnews.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"nbcnews.com", "www.nbcnews.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"buzzfeed.com", "www.buzzfeed.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://news.yahoo.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"news.yahoo.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.cnn.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"cnn.com", "www.cnn.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.espn.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"espn.com", "www.espn.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.huffpost.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"huffpost.com", "www.huffpost.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.foxnews.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"foxnews.com", "www.foxnews.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.npr.org/sections/news/", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"npr.org", "www.npr.org"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.msnbc.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"msnbc.com", "www.msnbc.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.nbcnews.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"nbcnews.com", "www.nbcnews.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.reuters.com", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"reuters.com", "www.reuters.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://news.google.com/topstories", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"news.google.com"}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////

	p.Searches = append(p.Searches, Search{SearchEngine_Bing, []string{"tornado last night", "tornado damage yesterday", "tornado last week near me"}, 1, 3})
	p.Searches = append(p.Searches, Search{SearchEngine_Yahoo, []string{"latest covid-19 news", "covid-19 update", "covid-19 news near me"}, 2, 5})
}

func ProfileScheduleFunc_NewsAddict(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("08:00")) || t.After(String2Time24("16:00")) {
		return false
	}

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

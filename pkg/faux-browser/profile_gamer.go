package faux_browser

import "time"

// Generate a profile of sites that a gamer would visit
func (p *Profile) GenerateProfile_Gamer() {
	p.ScheduleFunc = ProfileScheduleFunc_Gamer
	p.PercentOfSiteRead = 80
	p.WPM = 200
	p.WPM_Jitter = 10

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"https://store.steampowered.com", SITE_CANCLICKLINK, []string{"steampowered.com", "store.steampowered.com"}, 60, 180})
	p.Sites = append(p.Sites, Site{"https://www.youtube.com/", SITE_CANCLICKLINK, []string{"youtube.com", "www.youtube.com"}, 300, 600})
	p.Sites = append(p.Sites, Site{"https://www.newworld.com/", SITE_DONTCLICKLINKS, []string{"newworld.com", "www.newworld.com"}, 300, 600})
	p.Sites = append(p.Sites, Site{"https://forums.newworld.com/", SITE_CANCLICKLINK, []string{"newworld.com", "www.newworld.com", "forums.newworld.com"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.minecraft.net/", SITE_CANCLICKLINK, []string{"minecraft.net", "www.minecraft.net"}, 0, 0})
	p.Sites = append(p.Sites, Site{"https://www.minecraft.net/en-us/community", SITE_CANCLICKLINK, []string{"minecraft.net", "www.minecraft.net"}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////

	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"when is half life 3 coming out", "half life 3", "half life 3 release date", "why won't gabe release half life 3"}, 1, 4})
	p.Searches = append(p.Searches, Search{SearchEngine_Bing, []string{"animal crossing", "latest animal crossing update", "when is the next animal crossing game coming out"}, 1, 3})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"chess games", "online chess games", "play chess online", "top 10 chess players"}, 1, 2})
	p.Searches = append(p.Searches, Search{SearchEngine_Bing, []string{"roblox", "latest roblox update", "roblox crashed", "make money from roblox", "is roblox better than minecraft"}, 1, 3})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"minecraft update", "minecraft crashed", "change minecraft skin", "best minecraft mods", "too many minecraft mods", "why is minecraft better than terraria", "future minecraft features"}, 1, 3})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"steam sale", "steam coupons", "when is the next steam sale", "steam christmas in july sale"}, 1, 2})
	p.Searches = append(p.Searches, Search{SearchEngine_Yahoo, []string{"latest new world patch", "can I play new world on Mac", "when is New World being released on stadia", "aeternum tips and tricks", "aeternum leveling tips", "new world cheats", "new world hack"}, 1, 3})
}

func ProfileScheduleFunc_Gamer(t time.Time) (atdesk bool) {
	// Is this outside of normal work hours?
	if t.Before(String2Time24("08:30")) || t.After(String2Time24("17:30")) {
		return false
	}

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

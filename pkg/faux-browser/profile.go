package faux_browser

import (
	"crypto/rand"
	"math/big"
	"os"
	"time"
)

type Profile struct {
	Browser           Browser             // Browser the user will use
	PercentOfSiteRead int                 // Used with to calculate the number of words for WPM/WPM_Jitter. Range is 0-100.
	Searches          []Search            // Searches the user will conduct
	ScheduleFunc      ProfileScheduleFunc // Used by the Browser to determine if the user is "at their desk"
	Sites             []Site              // Sites the user will navigate to
	WPM               int                 // The number of words per minute the user can "read"
	WPM_Jitter        int                 // Jitter to the user's WPM
}

// Used by Browser to see if a profile should be clicking the link, based on
// the time provided. Used to simulate users operating within certain windows
// of time (i.e., business hours), and days of the week.
type ProfileScheduleFunc func(time.Time) bool

func String2Time24(s string) (t time.Time) {
	t, _ = time.Parse("15:04", s)
	return
}

func NewProfile() (p *Profile) {
	p = &Profile{}
	p.Browser = *NewBrowser()

	// Generate a profile
	// Check if the environment variable is set, otherwise "weighted random"
	switch os.Getenv("FAUX_B_ROLE") {
	case "cybersec":
		p.GenerateProfile_Cybersecurity()
	case "dev":
		p.GenerateProfile_Developer()
	case "driven":
		p.GenerateProfile_Employee_Driven()
	case "employee":
		p.GenerateProfile_Employee_Average()
	case "gamer":
		p.GenerateProfile_Gamer()
	case "manage":
		p.GenerateProfile_Management()
	case "news":
		p.GenerateProfile_NewsAddict()
	case "reddit":
		p.GenerateProfile_Redditor()
	case "research":
		p.GenerateProfile_Researcher()
	case "sales":
		p.GenerateProfile_Sales()
	case "tech":
		p.GenerateProfile_IT()
	default:
		if FakeWeightedRandomCheck(30) {
			p.GenerateProfile_Management()
		} else if FakeWeightedRandomCheck(20) {
			p.GenerateProfile_Employee_Driven()
		} else if FakeWeightedRandomCheck(10) {
			p.GenerateProfile_Sales()
		} else if FakeWeightedRandomCheck(10) {
			p.GenerateProfile_Redditor()
		} else if FakeWeightedRandomCheck(10) {
			p.GenerateProfile_IT()
		} else if FakeWeightedRandomCheck(10) {
			p.GenerateProfile_NewsAddict()
		} else if FakeWeightedRandomCheck(10) {
			p.GenerateProfile_Developer()
		} else if FakeWeightedRandomCheck(10) {
			p.GenerateProfile_Cybersecurity()
		} else if FakeWeightedRandomCheck(10) {
			p.GenerateProfile_Gamer()
		} else if FakeWeightedRandomCheck(30) {
			p.GenerateProfile_Threat()
		} else if FakeWeightedRandomCheck(20) {
			p.GenerateProfile_Researcher()
		} else {
			p.GenerateProfile_Employee_Average()
		}
	}

	// Add the common sites to all profiles
	p.GenerateProfile_Common()

	return
}

func FakeWeightedRandomCheck(check int64) bool {
	x, _ := rand.Int(rand.Reader, big.NewInt(check))
	return x.Int64() == 1
}

///////////////////////////////////////////////////////////////////////////////

func (p *Profile) StartBrowsing() {
	var sleeptime int64
	var urls []string
	var urlqueue []string

	for {
		// Reset sleeptime
		sleeptime = 30

		// Check if we are in the schedule to be browsing
		if p.ScheduleFunc(time.Now()) {
			if FakeWeightedRandomCheck(1000) {
				// Very small chance of ignoring the queue and visiting a site
				// from the list of browser.navigatabledomains
				i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(p.Browser.NavigatableDomains))))
				u := "https://" + p.Browser.NavigatableDomains[i.Int64()]

				// Visit the random site!
				urls, sleeptime = p.Browser.VisitSite(Site{u, SITE_CANCLICKANYLINK, []string{}, 0, 0}, p)
				urlqueue = append(urlqueue, urls...)
			} else if len(urlqueue) > 0 {
				// If we have URLs queued up, then pick one at random to visit,
				// otherwise, visit a Site or perform a Search

				// Pick a url at random from the queue
				i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(urlqueue))))
				u := urlqueue[i.Int64()]

				// Remove it from the queue
				urlqueue = append(urlqueue[:i.Int64()], urlqueue[i.Int64()+1:]...)

				// Visit it and add any additional urls to the queue
				urls, sleeptime = p.Browser.VisitSite(Site{u, SITE_CANCLICKANYLINK, []string{}, 0, 0}, p)
				urlqueue = append(urlqueue, urls...)
			} else {
				// Are we searching or are we visiting a site directly?
				if FakeWeightedRandomCheck(30) {
					if len(p.Searches) > 0 {
						// Use a search engine and add a few items to the queue!
						i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(p.Searches))))
						urls, sleeptime = p.Browser.ConductSearch(p.Searches[i.Int64()], p)
						urlqueue = append(urlqueue, urls...)
					}
				} else {
					if len(p.Sites) > 0 {
						// Visit a site directly!
						i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(p.Sites))))
						urls, sleeptime = p.Browser.VisitSite(p.Sites[i.Int64()], p)
						urlqueue = append(urlqueue, urls...)
					}
				}
			}
		} else {
			// Wait 10-30 minutes if we are not in our schedule
			sleeptime = GetMinMax(6000, 18000)
		}

		// Sanity check of sleeptime to avoid issues
		if sleeptime < 30 {
			sleeptime = GetMinMax(15, 45)
		}

		// Sleep!
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}
}

func (p *Profile) GetWPM() int64 {
	max := int64(p.WPM + p.WPM_Jitter)
	min := int64(p.WPM - p.WPM_Jitter)
	wpm, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err == nil {
		return wpm.Int64() + min
	}
	return 200
}

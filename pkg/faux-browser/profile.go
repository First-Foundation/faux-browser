package faux_browser

import (
	"crypto/rand"
	"math/big"
	"os"
)

type Profile struct {
	Browser           Browser  // Browser the user will use
	PercentOfSiteRead int      // Used with to calculate the number of words for WPM/WPM_Jitter. Range is 0-100.
	Searches          []Search // Searches the user will conduct
	Sites             []Site   // Sites the user will navigate to
	WPM               int      // The number of words per minute the user can "read"
	WPM_Jitter        int      // Jitter to the user's WPM
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
		if FakeWeightedRandomCheck(100) {
			p.GenerateProfile_Management()
		} else if FakeWeightedRandomCheck(100) {
			p.GenerateProfile_Employee_Driven()
		} else if FakeWeightedRandomCheck(50) {
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
		} else if FakeWeightedRandomCheck(100) {
			p.GenerateProfile_Threat()
		} else if FakeWeightedRandomCheck(50) {
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
	return x == big.NewInt(1)
}

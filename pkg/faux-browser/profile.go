package faux_browser

import (
	"crypto/rand"
	"math/big"
	"os"
)

type Profile struct {
	Browser    Browser // Browser the user will use
	Sites      []Site  // Sites the user will navigate to
	WPM        int     // The number of words per minute the user can "read"
	WPM_Jitter int     // Jitter to the user's WPM
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

	return
}

func FakeWeightedRandomCheck(check int64) bool {
	x, _ := rand.Int(rand.Reader, big.NewInt(check))
	return x == big.NewInt(1)
}

///////////////////////////////////////////////////////////////////////////////
// ROLE SPECIFIC PROFILES
///////////////////////////////////////////////////////////////////////////////

// Generate a profile of sites that someone in Cybersecurity would visit
func (p *Profile) GenerateProfile_Cybersecurity() {
	p.WPM = 250
	p.WPM_Jitter = 50

	p.Sites = append(p.Sites, Site{"https://book.hacktricks.xyz", true, false, []string{"book.hacktricks.xyz"}})
	p.Sites = append(p.Sites, Site{"https://www.virustotal.com", true, false, []string{"virustotal.com", "www.virustotal.com"}})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", false, false, []string{}})
}

// Generate a profile of sites that a Developer might visit
func (p *Profile) GenerateProfile_Developer() {
	p.WPM = 200
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://docs.microsoft.com/en-us/", true, false, []string{"docs.microsoft.com"}})
	p.Sites = append(p.Sites, Site{"https://stackoverflow.com/questions/tagged/c%2B%2B", true, false, []string{"stackoverflow.com"}})
}

// Generate a profile of sites that an average employee would visit
func (p *Profile) GenerateProfile_Employee_Average() {
	p.WPM = 200
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://www.facebook.com/", true, false, []string{"facebook.com", "www.facebook.com"}})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com/", true, false, []string{"linkedin.com", "www.linkedin.com"}})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", true, false, []string{"youtube.com", "www.youtube.com"}})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", true, false, []string{"buzzfeed.com", "www.buzzfeed.com"}})

	// USER TODO: Add company related sites!
}

// Generate a profile of sites that a "driven" employee would visit
func (p *Profile) GenerateProfile_Employee_Driven() {
	p.WPM = 300
	p.WPM_Jitter = 50

	p.Sites = append(p.Sites, Site{"https://www.themuse.com/advice/47-habits-of-highly-successful-employees", false, false, []string{}})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", true, false, []string{"youtube.com", "www.youtube.com"}})

	// USER TODO: Add company related sites!
}

// Generate a profile of sites that a gamer would visit
func (p *Profile) GenerateProfile_Gamer() {
	p.WPM = 200
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://store.steampowered.com", true, false, []string{"steampowered.com", "store.steampowered.com"}})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", true, false, []string{"youtube.com", "www.youtube.com"}})
}

// Generate a profile of sites that someone in I.T. would visit
func (p *Profile) GenerateProfile_IT() {
	p.WPM = 200
	p.WPM_Jitter = 30

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/InformationTechnology/", true, false, []string{"reddit.com", "www.reddit.com"}})
}

// Generate a profile of sites that a CEO or department lead would visit
func (p *Profile) GenerateProfile_Management() {
	p.WPM = 300
	p.WPM_Jitter = 10

	p.Sites = append(p.Sites, Site{"https://www.foreignaffairs.com", true, false, []string{"foreignaffairs.com", "www.foreignaffairs.com"}})
	p.Sites = append(p.Sites, Site{"https://www.ted.com", true, false, []string{"ted.com", "www.ted.com"}})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com", true, false, []string{"linkedin.com", "www.linkedin.com"}})
	p.Sites = append(p.Sites, Site{"https://www.ceo.com", true, false, []string{"ceo.com", "www.ceo.com"}})
	p.Sites = append(p.Sites, Site{"http://wsj.com/", true, false, []string{"wsj.com", "www.wsj.com"}})
}

// Generate a profile of someone who browsers news websites
func (p *Profile) GenerateProfile_NewsAddict() {
	p.WPM = 200
	p.WPM_Jitter = 20

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/news", true, true, []string{}})
	p.Sites = append(p.Sites, Site{"https://www.nbcnews.com", true, false, []string{"nbcnews.com", "www.nbcnews.com"}})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", true, false, []string{"buzzfeed.com", "www.buzzfeed.com"}})
}

// Generate a profile of someone who browses reddit at work all day
func (p *Profile) GenerateProfile_Redditor() {
	p.WPM = 150
	p.WPM_Jitter = 50

	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/AskReddit", true, false, []string{"reddit.com", "www.reddit.com"}})
	p.Sites = append(p.Sites, Site{"https://www.reddit.com/r/pics", true, false, []string{"reddit.com", "www.reddit.com"}})
	p.Sites = append(p.Sites, Site{"https://www.buzzfeed.com", true, false, []string{"buzzfeed.com", "www.buzzfeed.com"}})
}

// Generate a profile for someone who likes to research
func (p *Profile) GenerateProfile_Researcher() {
	p.WPM = 250
	p.WPM_Jitter = 25

	p.Sites = append(p.Sites, Site{"https://en.wikipedia.org", true, false, []string{"en.wikipedia.org"}})
}

// Generate a profile of sites a sales person would visit
func (p *Profile) GenerateProfile_Sales() {
	p.WPM = 225
	p.WPM_Jitter = 25

	p.Sites = append(p.Sites, Site{"https://www.salesforce.com", true, false, []string{"salesforce.com", "www.salesforce.com"}})
	p.Sites = append(p.Sites, Site{"https://www.neilpatel.com", true, false, []string{"neilpatel.com", "www.neilpatel.com"}})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com/", true, false, []string{"linkedin.com", "www.linkedin.com"}})
}

// Generate a profile of sites that an "unintentional insider threat" would visit
func (p *Profile) GenerateProfile_Threat() {
	p.WPM = 150
	p.WPM_Jitter = 25

	// USER TODO: Modify the below to match your needs!
	p.Sites = append(p.Sites, Site{"https://zh.wikipedia.org/", true, false, []string{"zh.wikipedia.org"}})
	p.Sites = append(p.Sites, Site{"https://ru.wikipedia.org", true, false, []string{"ru.wikipedia.org"}})
	p.Sites = append(p.Sites, Site{"https://gchq.github.io/CyberChef", false, false, []string{}})
	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", true, false, []string{"youtube.com", "www.youtube.com"}})
	p.Sites = append(p.Sites, Site{"http://www.reddit.com/", true, false, []string{"reddit.com", "www.reddit.com"}})
}

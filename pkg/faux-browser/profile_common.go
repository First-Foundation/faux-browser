package faux_browser

// Add common sites that everyone will be visiting
func (p *Profile) GenerateProfile_Common() {
	// DON'T MESS WITH ANYTHING OTHER THAN ADDING TO SITES AND SEARCHES

	///////////////////////////////////////////////////////////////////////////
	// SITES TO VISIT                                                        //
	///////////////////////////////////////////////////////////////////////////

	p.Sites = append(p.Sites, Site{"http://www.youtube.com/", SITE_CANCLICKLINK, []string{"youtube.com", "www.youtube.com"}, 60, 300})
	p.Sites = append(p.Sites, Site{"https://www.linkedin.com/", SITE_CANCLICKLINK | SITE_CANCLICKANYLINK, []string{"linkedin.com", "www.linkedin.com"}, 300, 0})
	p.Sites = append(p.Sites, Site{"https://www.facebook.com/", SITE_CANCLICKLINK, []string{"facebook.com", "www.facebook.com"}, 60, 180})
	p.Sites = append(p.Sites, Site{"https://trends.google.com/trends/trendingsearches/daily?geo=US", SITE_CANCLICKANYLINK, []string{}, 0, 0})

	///////////////////////////////////////////////////////////////////////////
	// SEARCHES TO CONDUCT                                                   //
	///////////////////////////////////////////////////////////////////////////

	// Shouldn't mix politics at work...
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"trump", "donald trump", "biden", "joe biden", "pelosi", "nancy pelosi", "mcconnell", "mitch mcconnell"}, 1, 1})

	// Covid 19 stuff
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"covid19", "coronavirus", "covid vaccination mandate", "covid mandate requirements", "covid testing near me"}, 1, 2})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"covid symptoms", "covid19 symptoms", "coronavirus symptoms", "cold symptoms", "cold flue or covid", "covid testing", "flu shot"}, 1, 1})

	// Results from googling for top sports teams
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"dallas cowboys", "new york yankees", "new york knicks", "barcelona", "real madrid", "golden state warriors", "los angeles lakers", "new england patriots"}, 1, 2})

	// Results from googling for celebrities
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"dwayne johnson", "kim kardashian", "cristiano ronaldo", "justin bieber", "oprah winfrey", "ariana grande", "rihanna", "angelina jolie", "jay-z", "howard stern", "ellen degeneres", "katy perry", "ben affleck", "jennifer lopez", "sean combs", "rush limbaugh", "jerry seinfeld", "drake", "jennifer aniston", "ryan reynolds", "donald trump", "simon cowell", "celine dion", "kevin hart", "rloyd mayweather", "adele", "will smith", "tiger woods"}, 1, 2})

	// Results from googling for 2021 movies
	p.Searches = append(p.Searches, Search{SearchEngine_Bing, []string{"raya and the last dragon", "judas and the black messiah", "our friend", "the croods", "halloween kills", "ron's gone wrong", "the last duel", "venom let there be carnage", "the green knight", "dune"}, 1, 2})

	// Results from googling for netflix originals
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"netflix you", "netflix maid", "locke and key", "squid game", "midnight mass netflix", "netflix legacies", "night teeth", "the sinner netflix", "on my block", "365 days show", "clickbait netflix", "inside job netflix show"}, 1, 2})

	// Results from googling top songs of 2021 (the artists)
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"lil nas x", "justin beiber", "jack harlow", "drumkoon", "olivia rodrigo", "bts", "ariana grande", "nightcrawlers", "ed sheeran", "camila cobello", "silk sonic", "doja cat", "the weeknd", "polo g"}, 1, 1})

	// Results from googling top songs of 2021
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"montero", "call me by your name", "stay", "industry baby", "music playlist", "good 4 u", "butter by bts", "save your tears", "friday dopamine reedit", "peaches justin beiber", "bad habits ed sheeran", "drivers license olivia rodrigo", "don't go yet song", "leave the door open by silk sonic", "doja cat kiss me more", "the weekend's take my breath", "rapstar polo g"}, 1, 1})

	// Results from googling "universities"
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"harvard", "stanford", "mit", "princeton", "yale", "columbia university", "caltech", "university of pennsylvania", "duke university", "berkley university", "university of chicago", "cornell university", "university of california", "university of michigan", "johns hopkins university", "brown university", "vanderbilt university", "rice university", "georgetown university", "emory university", "wgu", "northwestern university", "university of virginia", "dartmouth college"}, 1, 1})

	// Results from googling "top books of 2021"
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"klara and the sun", "harlem shuffle", "the other black girl", "beautiful world where are you", "the four winds", "malibu rising", "somebody's daughter", "detransition baby", "crying in h mart", "one last stop book", "second place", "aftershocks memoir", "no one is talking about this", "the last thing he told me", "filthy animals", "let me tell you what i mean", "cloud cuckoo land", "empire of pain", "matrix", "people we meet on vacation", "infinite country", "great circle", "a slow fire burning", "project hail mary", "while justice sleeps", "caul baby a novel", "of women and salt"}, 1, 1})

	// Pulled top 50 from top 500 restaurants
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"mcdonalds", "starbucks", "chickfila", "taco bell", "wendys", "burger king", "dunkin donuts", "subway", "dominos", "chipotle", "sonic drive-in", "pizza hut", "panera bread", "kfc", "popeyes", "arbys", "dairy queen", "little ceasars", "panda express", "jack in the box", "olive garden", "papa johns", "buffalo wild wings", "applebees", "chili's", "whataburger", "texas roadhouse", "ihop", "outback stakhouse", "zaxby's", "culver's", "jimmy john's", "wingstop", "hardees", "cracker barrel", "red lobster", "denny's", "raising cane's", "five guys", "jersey mike's", "longhorn steakhouse", "cheesecake factory", "bojangles", "carl's jr", "in-n-out", "red robin", "el pollo loco", "del taco", "firehouse subs", "qdoba"}, 1, 2})

	// Office supplies!
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"paper", "printer paper", "three-hole punched paper", "graph paper", "tracing paper", "carbon paper", "color card stock", "heavy-duty card stock", "wrapping paper", "greeting cards", "business cards", "letterhead", "envelopes", "#10 regular envelopes", "legal envelopes", "padded legal envelope mailers", "postage stamps", "envelope sealer", "packaging bubble"}, 1, 2})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"boxes", "cardboard", "cardboard boxes", "small cardboard boxes", "medium cardboard boxes", "large cardboard boxes", "extra large cardboard boxes", "notebooks", "notepads", "moleskine", "composition notebooks", "spiral-bound notebooks", "legal pads", "steno pads", "sticky notes", "bookends", "paperweight", "magazine holders", "bulletin board", "pushpins"}, 1, 2})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"binders", "binder tabs", "binder pockets", "binder document holders", "clear binder document holders", "hole puncher", "three-hole puncher", "filing cabinet", "manila folders", "folders", "hanging folders", "folder tabs", "t-square", "compass", "protractor", "ruler", "erase board", "erase marker", "erase spray", "white out", "letter opener", "pen holder", "supply tray"}, 1, 2})
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"office supplies", "small office supplies", "staplers", "staples", "stapler removers", "scissors", "box cutters", "paperclips", "binder clips", "cellophane tape", "masking tape", "packing tape", "duct tape", "bookmarks", "white glue", "sticky notes", "bookmark sticky flags", "rubber cement", "tacky wall mount gum", "hanging hooks", "magnifying glasses", "pencils", "pencil sharpener", "mechanical pencils", "mechanical pencil lead refills", "erasers", "pens", "all-purpose markers", "highlighters", "rubber stamps", "ink pad"}, 1, 2})

	// 2021 Cellphones
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"iphone 13 pro max", "samsung galaxy s21 ultra", "iphone 13", "google pixel 6", "google pixel 5a", "iphone 13 pro", "oneplus 9 pro", "samsung galaxy s21", "samsung galaxy z fold 3", "iphone 11", "moto g power", "iphone se"}, 1, 2})

	// Pulled from top 100 classic movies, according to rotten tomatoes
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"it happened one night", "modern times", "wizard of oz", "citizen kane", "casablanca", "all about eve", "the philadelphia story", "sunset boulevard", "the cabinet of dr. caligari", "rebecca", "his girl friday", "a night at the opera", "the third man", "rear window", "seven samurai", "singin in the rain", "la grande illusion", "all quiet on the western front", "snow white and the seven dwarfs", "an american in paris", "the kid", "on the waterfront", "the best years of our lives", "metropolis", "the adventures of robin hood", "north by northwest", "laura", "king kong", "nosferatu", "shadow of a doubt", "psycho", "a hard day's night", "top hat", "the bride of frankenstein", "12 angry men", "marty", "the treasure of sierra madre", "the lady eve", "the lady vanishes", "chinatown", "dr strangelove", "frankenstein", "the french connection", "the thin man", "the 39 steps", "the gold rush", "kind hearts and coronets", "battleship potemkin", "touch of evil", "a streetcar named desire"}, 1, 3})

	// "Most Populer TV On Rotten Tomatoes" - 2021.OCT.29
	p.Searches = append(p.Searches, Search{SearchEngine_Google, []string{"you season 3", "squid game season 1", "only murders in the building season 1", "locke & key season 2", "maid season 1", "midnight mass", "chucky season 1", "invasion season 1", "succession season 3", "the sinner season 4", "foundation season 1", "inside job season 1", "money heist part 5", "dopesick season 1", "another life season 2", "curb your enthusiasm", "the chestnut man season 1", "sex education season 3", "ted lasso season 2", "my name season 1"}, 1, 2})
}

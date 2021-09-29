package faux_browser

type Search struct {
	Engine            SearchEngine
	Queries           []string
	MinResultsToClick int
	MaxResultsToClick int
}

type SearchEngine byte

const (
	SearchEngine_Google SearchEngine = iota + 1
	SearchEngine_Yahoo
	SearchEngine_Bing
)

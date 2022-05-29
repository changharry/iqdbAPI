package iqdbAPI

type Response struct {
	Success bool
	Results []Result
}

type Result struct {
	Head       string
	Url        string
	Titles     string
	Height     string
	Width      string
	Category   string
	Similarity string
}

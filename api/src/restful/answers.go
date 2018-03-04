package restful

import "wuzzapcom/Coursework/api/src/bibtex"

type Error struct {
	Message string `json:"errorMessage"`
}

type Search struct {
	Results bibtex.Items `json:"results"`
}

type Success struct {
	Message string `json:"successMessage"`
}

func (search Search) String() string {

	return search.Results.String()

}

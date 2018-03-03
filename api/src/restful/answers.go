package restful

import "wuzzapcom/Coursework/api/src/bibtex"

type Error struct {
	Message string `json:"errorMessage"`
}

type Search struct {
	Results bibtex.Items `json:"results"`
}

func (search Search) String() string {

	return search.Results.String()

}

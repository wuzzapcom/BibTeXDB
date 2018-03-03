package fetchers

import (
	"fmt"
	"strings"
)

const yearStringLength = 4

//IndustryIdentifier ..
type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type volumeInfo struct {
	Title               string               `json:"title"`
	Authors             []string             `json:"authors"`
	Description         string               `json:"description"`
	Publisher           string               `json:"publisher"`
	Date                string               `json:"publishedDate"`
	Language            string               `json:"language"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
}

func (v *volumeInfo) getBibtexAuthors() (result string) {

	for i, author := range v.Authors {
		if i == 0 {
			result += author
		} else {
			result += " and " + author
		}
	}
	return result

}

func (v *volumeInfo) getBibtexISBN() string {

	for _, ident := range v.IndustryIdentifiers {
		if ident.Type == "ISBN_13" {
			return ident.Identifier
		}
	}

	return ""
}

func (v *volumeInfo) getBibtexYear() string {
	splittedDate := strings.Split(v.Date, "-")
	for _, year := range splittedDate {
		if len(year) == yearStringLength {
			return year
		}
	}
	return ""
}

type items struct {
	ID         string     `json:"id"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo volumeInfo `json:"volumeInfo"`
}

func (i items) String() (result string) {
	return fmt.Sprintf("\n\t\tID: %s\n\t\tSelfLink: %s\n\t\tVolumeInfo: %s", i.ID, i.SelfLink, i.VolumeInfo)
}

type mainGoogleAPIResponse struct {
	TotalItems int     `json:"totalItems"`
	Items      []items `json:"items"`
}

func (m mainGoogleAPIResponse) String() (result string) {
	result = fmt.Sprintf("Total items: %d\nItems:\n", m.TotalItems)
	for _, item := range m.Items {
		result += fmt.Sprintf("\t[%s]\n", item)
	}
	return result
}

package fetchers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"wuzzapcom.io/Coursework/pkg/bibtex"
)

//GoogleFetcher ..
type GoogleFetcher struct {
	APIToken string
}

//FetchWithTitleAndAuthor ..
func (fetcher *GoogleFetcher) FetchWithTitleAndAuthor(title string, author string) ([]bibtex.Item, error) {

	if fetcher.APIToken == "" {
		return nil, errors.New("Provided no API Token for Google Books")
	}

	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s+inauthor:%s&key=%s",
		url.QueryEscape(title),
		url.QueryEscape(author),
		url.QueryEscape(fetcher.APIToken),
	)

	return fetcher.fetch(url)

}

//FetchWithString ..
func (fetcher *GoogleFetcher) FetchWithString(text string) ([]bibtex.Item, error) {

	if fetcher.APIToken == "" {
		return nil, errors.New("Provided no API Token for Google Books")
	}

	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&key=%s",
		url.QueryEscape(text),
		url.QueryEscape(fetcher.APIToken),
	)

	return fetcher.fetch(url)
}

func (fetcher *GoogleFetcher) fetch(url string) ([]bibtex.Item, error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var googleResponse mainGoogleAPIResponse

	json.Unmarshal(body, &googleResponse)

	fmt.Println(googleResponse)

	var result []bibtex.Item

	for _, item := range googleResponse.Items {

		bibTex := bibtex.Item{
			Ident:     "empty",
			Title:     item.VolumeInfo.Title,
			Author:    item.VolumeInfo.getBibtexAuthors(),
			Publisher: item.VolumeInfo.Publisher,
			Year:      item.VolumeInfo.getBibtexYear(),
			Language:  item.VolumeInfo.Language,
			ISBN:      item.VolumeInfo.getBibtexISBN(),
			URL:       item.SelfLink,
		}

		result = append(result, bibTex)

	}

	return result, nil

}

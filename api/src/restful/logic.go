package restful

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"wuzzapcom/Coursework/api/src/bibtex"
)

func searchCheckInput(w http.ResponseWriter, r *http.Request) (url.Values, error){
	parameters, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		internalServerError(w, "Parsing parameter error")
		return nil, err
	}

	if len(parameters["request"]) == 0 {
		answer, err := json.Marshal(Error{"No request provided"})
		if err != nil {
			internalServerError(w, "Internal server error")
			return nil, err
		}

		w.WriteHeader(500)
		w.Write(answer)
	}
	return parameters, nil
}

func search(w http.ResponseWriter, values url.Values){
	var result bibtex.Items
	for _, request := range values["request"] {
		res, err := fetcher.FetchWithString(request)
		if err != nil{
			fmt.Println(err)
			internalServerError(w, "Parsing parameter error")
			return
		}
		result.Append(res)
	}

	answer, err := json.Marshal(Search{result})
	if err != nil {
		fmt.Println(err)
		internalServerError(w, "Internal server error")
		return
	}

	w.WriteHeader(200)
	w.Write(answer)
}

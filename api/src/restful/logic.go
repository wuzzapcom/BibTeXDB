package restful

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"wuzzapcom/Coursework/api/src/bibtex"
	"io/ioutil"
	"wuzzapcom/Coursework/api/src/database"
)

func searchCheckInput(w http.ResponseWriter, r *http.Request) (url.Values, error){
	parameters, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		returnError(w, 500, "Parsing parameter error")
		return nil, err
	}

	if len(parameters["request"]) == 0 {
		answer, err := json.Marshal(Error{"No request provided"})
		if err != nil {
			returnError(w, 500, "Internal server error")
			return nil, err
		}

		w.WriteHeader(400)
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
			returnError(w, 500, "Parsing parameter error")
			return
		}
		result.Append(res)
	}

	answer, err := json.Marshal(Search{result})
	if err != nil {
		fmt.Println(err)
		returnError(w, 500, "Internal server error")
		return
	}

	w.WriteHeader(200)
	w.Write(answer)
}

func addBookCheckInput(w http.ResponseWriter, r *http.Request) ([]byte, error){
	body := r.Body
	//if  != nil {
	//	returnError(w, 400, "No JSON provided")
	//	return nil, err
	//}

	answer, err := ioutil.ReadAll(body)
	if err != nil {
		returnError(w, 400, "No JSON provided")
		return nil, err
	}

	return answer, nil

}

func addBook(w http.ResponseWriter, body []byte){
	var addingBooks bibtex.Item

	err := json.Unmarshal(body, &addingBooks)
	if err != nil{
		fmt.Println(err)
		returnError(w, 400, "Wrong JSON input")
		return
	}

	mongo := database.Mongo{}
	mongo.Connect()
	defer mongo.Disconnect()

	err = mongo.InsertTextbook(addingBooks)
	if err != nil {
		fmt.Println(err)
		returnError(w, 500, "Internal server error")
		return
	}

	answer, err := json.Marshal(Success{"OK"})
	if err != nil {
		fmt.Println(err)
		returnError(w, 500, "Internal server error")
		return
	}

	w.WriteHeader(200)
	w.Write(answer)

}

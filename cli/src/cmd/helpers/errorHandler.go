package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"wuzzapcom/Coursework/api/src/restful"
)

// ServerURL ..
const ServerURL = "http://localhost:8080/"

//Printable ..
type Printable interface {
	String() string
}

// HandleError ..
func HandleError(answer []byte) {
	var errorAnswer restful.Error
	err := json.Unmarshal(answer, &errorAnswer)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	fmt.Println(errorAnswer.Message)
}

// GetFromServer ..
func GetFromServer(url string, strct interface{}) (interface{}, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		HandleError(data)
		return nil, err
	}

	err = json.Unmarshal(data, &strct)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return nil, err
	}

	return strct, nil
}

// PrintResult ..
func PrintResult(printable Printable, outputFile string) {
	if outputFile == "" {
		fmt.Println(printable.String())
	} else {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			return
		}
		file.Write([]byte(printable.String()))
	}
}

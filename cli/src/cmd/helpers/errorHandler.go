package helpers

import (
	"bytes"
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

// LoadFromFileAndValidate ..
func LoadFromFileAndValidate(strct interface{}, inputFile string) ([]byte, error) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &strct)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return nil, err
	}

	return data, nil
}

// SendDataToServer ..
func SendDataToServer(data []byte, url string) error {
	response, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	answer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		HandleError(answer)
		return nil
	}

	var success restful.Success
	err = json.Unmarshal(answer, &success)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}
	return nil
}

// DeleteFile ..
func DeleteFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("Не удалось удалить файл", err)
	}
}

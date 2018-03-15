package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var getLecturersCommand = &cobra.Command{
	Use: "getLecturers",
	Run: getLecturers,
}

func getLecturers(cmd *cobra.Command, args []string) {

	url := "http://localhost:8080/getLecturers"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	if response.StatusCode != 200 {
		handleError(data)
		return
	}
	var lecturers restful.Lecturers
	err = json.Unmarshal(data, &lecturers)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	for _, lecturer := range lecturers.LecturerList {
		fmt.Println(lecturer.String())
	}

}

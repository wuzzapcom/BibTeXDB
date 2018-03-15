package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var addLecturerCommand = &cobra.Command{
	Use: "addLecturer",
	Run: addLecturer,
}

const inputFileLecturerFlag = "inputFile"

func addLecturer(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileFlag).Value.String()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var items common.Lecturer
	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := "http://localhost:8080/addLecturer"
	response, err := http.Post(url, "application/json", bytes.NewReader(data))

	answer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	if response.StatusCode != 200 {
		handleError(answer)
		return
	}

	var success restful.Success
	err = json.Unmarshal(answer, &success)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	fmt.Println(success)
}

func init() {
	addLecturerCommand.Flags().String(
		inputFileFlag,
		lecturerDefaultPath,
		"Set input file for lecturer",
	)
}

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/api/src/restful"

	"github.com/spf13/cobra"
)

var addLecturerCommand = &cobra.Command{
	Use:   "addLecturer",
	Run:   addLecturer,
	Short: "Отправить лектора на сервер из файла, заданного флагом.",
}

const inputFileLecturerFlag = "inputFile"

func addLecturer(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileLecturerFlag).Value.String()
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
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

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

	err = os.Remove(inputFile)
	if err != nil {
		fmt.Println("Не удалось удалить файл", err)
	}
}

func init() {
	addLecturerCommand.Flags().String(
		inputFileLecturerFlag,
		lecturerDefaultPath,
		"Set input file for lecturer",
	)
}

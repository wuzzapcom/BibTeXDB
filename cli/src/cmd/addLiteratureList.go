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

var addLiteratureListCommand = &cobra.Command{
	Use:   "addLiteratureList",
	Run:   addLiteratureList,
	Short: "Добавить список литературы из файла, заданного флагом.",
	Long: `Добавить книгу в список литературы. 
	Поле Year определяет, за какой год создается список литературы.
	Поля CourseTitle и DepartmentTitle определяют учебный курс, для которого создается список литературы.`,
}

const inputFileLiteratureListFlag = "inputFile"

func addLiteratureList(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileLiteratureListFlag).Value.String()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var items common.LiteratureList
	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := "http://localhost:8080/addLiteratureList"
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
	addLiteratureListCommand.Flags().String(
		inputFileLiteratureListFlag,
		literatureListDefaultPath,
		"Set input file for literatureList",
	)
}

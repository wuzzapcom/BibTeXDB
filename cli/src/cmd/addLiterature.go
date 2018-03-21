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

var addLiteratureCommand = &cobra.Command{
	Use:   "addLiterature",
	Run:   addLiterature,
	Short: "Добавить книгу в список литературы из файла, заданного флагом.",
	Long: `Добавить книгу в список литературы. 
	Поле BookIdent задает идентификатор книги в формате BibTeX.
	Поле Year определяет, за какой год используется список литературы.
	Поля CourseTitle и DepartmentTitle определяют учебный курс.
	`,
}

const inputFileLiteratureFlag = "inputFile"

func addLiterature(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileLiteratureFlag).Value.String()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var items common.Literature
	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := "http://localhost:8080/addLiterature"
	response, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}

	answer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	if response.StatusCode != 200 {
		handleError(answer)
		return
	}

	var success restful.Success
	err = json.Unmarshal(answer, &success)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	fmt.Println(success)

	err = os.Remove(inputFile)
	if err != nil {
		fmt.Println("Не удалось удалить файл", err)
	}
}

func init() {
	addLiteratureCommand.Flags().String(
		inputFileLiteratureFlag,
		literatureDefaultPath,
		"Set input file for literature",
	)
}

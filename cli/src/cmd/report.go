package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var reportCommand = &cobra.Command{
	Use:   "report",
	Run:   report,
	Short: "Выполнить создание списка литературы.",
}

const inputFileReportFlag = "inputFile"
const outputFileReportFlag = "outputFile"

func report(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileReportFlag).Value.String()
	outputFileName := cmd.Flag(outputFileReportFlag).Value.String()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var list common.LiteratureList
	err = json.Unmarshal(data, &list)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := "http://localhost:8080/generateReport"
	response, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
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

	if outputFileName == "" {
		fmt.Println(string(answer))
	} else {
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			handleError(answer)
			return
		}
		outputFile.Write(answer)
	}

	// err = os.Remove(inputFile)
	// if err != nil {
	// 	fmt.Println("Не удалось удалить файл", err)
	// 	return
	// }
}

func init() {
	reportCommand.Flags().String(
		inputFileReportFlag,
		literatureListDefaultPath,
		"Set input file for migrate",
	)
	reportCommand.Flags().String(
		outputFileReportFlag,
		"",
		"Название файла, куда будет сгенерирован отчет(с расширением .bib). Без указания будет выведен текстом.",
	)
}

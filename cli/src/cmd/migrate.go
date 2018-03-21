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

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Run:   migrate,
	Short: "Выполнить копирование списка литературы с одного учебного года на другой.",
}

const inputFileMigrateFlag = "inputFile"

func migrate(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileMigrateFlag).Value.String()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	var migrate common.Migrate
	err = json.Unmarshal(data, &migrate)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := "http://localhost:8080/migrateLiteratureList"
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
	migrateCommand.Flags().String(
		inputFileMigrateFlag,
		migrateDefaultPath,
		"Set input file for migrate",
	)
}

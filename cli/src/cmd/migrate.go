package cmd

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/common"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Run:   migrate,
	Short: "Выполнить копирование списка литературы с одного учебного года на другой.",
	Long: `Выполнить копирование списка литературы с одного учебного года на другой. 
	При миграции список литературы, куда будет выполнено копирование, должен существовать и быть пустым.`,
}

const inputFileMigrateFlag = "inputFile"

func migrate(cmd *cobra.Command, args []string) {
	inputFile := cmd.Flag(inputFileMigrateFlag).Value.String()
	var migrate common.Migrate
	data, err := helpers.LoadFromFileAndValidate(migrate, inputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	url := helpers.ServerURL + "migrateLiteratureList"
	err = helpers.SendDataToServer(data, url)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	helpers.DeleteFile(inputFile)
}

func init() {
	migrateCommand.Flags().String(
		inputFileMigrateFlag,
		migrateDefaultPath,
		"Set input file for migrate",
	)
}

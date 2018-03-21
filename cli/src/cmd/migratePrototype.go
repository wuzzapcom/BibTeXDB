package cmd

import (
	"fmt"
	"os"
	"wuzzapcom/Coursework/api/src/common"

	"github.com/spf13/cobra"
)

var migratePrototypeCommand = &cobra.Command{
	Use:   "migratePrototype",
	Run:   migratePrototype,
	Short: "Получить заготовку JSON для миграции в файл, определяемый флагом.",
	Long:  "Получить заготовку JSON для миграции. После чего следует заполнить его вручную и отправить соответствующей командой.",
}

const migrateOutputFile = "outputFile"
const migrateDefaultPath = "migrate.txt"

func migratePrototype(cmd *cobra.Command, args []string) {
	outputFile := cmd.Flag(migrateOutputFile).Value.String()

	resultFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	resultFile.WriteString(common.GetMigrateExample().String())
	resultFile.Close()

	fmt.Println(fmt.Sprintf("Open %s and fill prototype struct with correct data", outputFile))
}

func init() {
	migratePrototypeCommand.Flags().String(
		migrateOutputFile,
		migrateDefaultPath,
		"Set output file for Migrate prototype",
	)
}

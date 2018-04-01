package departments

import "github.com/spf13/cobra"

// DepartmentCommand ..
var DepartmentCommand = &cobra.Command{
	Use:   "department",
	Short: "Команда для работы с курсами",
}

func init() {
	DepartmentCommand.AddCommand(addCommand)
	DepartmentCommand.AddCommand(getCommand)
	DepartmentCommand.AddCommand(prototypeCommand)
}

package literature

import "github.com/spf13/cobra"

// LiteratureCommand ..
var LiteratureCommand = &cobra.Command{
	Use:   "literature",
	Short: "Команда для работы с литературой",
}

func init() {
	LiteratureCommand.AddCommand(addCommand)
	LiteratureCommand.AddCommand(getCommand)
	LiteratureCommand.AddCommand(prototypeCommand)
}

package books

import "github.com/spf13/cobra"

// BookCommand ..
var BookCommand = &cobra.Command{
	Use:   "book",
	Short: "Команда для работы с книгами",
}

func init() {
	BookCommand.AddCommand(addCommand)
	BookCommand.AddCommand(getCommand)
	BookCommand.AddCommand(prototypeCommand)
}

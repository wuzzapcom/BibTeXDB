package literature_lists

import "github.com/spf13/cobra"

// LiteratureListCommand ..
var LiteratureListCommand = &cobra.Command{
	Use:   "literatureList",
	Short: "Команда для работы с курсами",
}

func init() {
	LiteratureListCommand.AddCommand(addCommand)
	LiteratureListCommand.AddCommand(getCommand)
	LiteratureListCommand.AddCommand(prototypeCommand)
}

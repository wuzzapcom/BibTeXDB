package lecturers

import "github.com/spf13/cobra"

// LecturerCommand ..
var LecturerCommand = &cobra.Command{
	Use:   "lecturer",
	Short: "Команда для работы с лекторами",
}

func init() {
	LecturerCommand.AddCommand(addCommand)
	LecturerCommand.AddCommand(getCommand)
	LecturerCommand.AddCommand(prototypeCommand)
}

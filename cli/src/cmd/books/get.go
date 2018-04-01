package books

import (
	"fmt"
	"wuzzapcom/Coursework/api/src/restful"
	"wuzzapcom/Coursework/cli/src/cmd/helpers"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Run:   getBooks,
	Short: "Получить список книг в формате BibTeX, сохраненных в базе данных.",
}

var getBooksOutputFlag = "toFile"

func getBooks(cmd *cobra.Command, args []string) {
	output := cmd.Flag(getBooksOutputFlag).Value.String()

	url := helpers.ServerURL + "getBooks"
	var books restful.Books

	answer, err := helpers.GetFromServer(url, &books)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}

	casted, ok := answer.(*restful.Books)
	if !ok {
		fmt.Println("Type cast error")
		return
	}
	books = *casted

	helpers.PrintResult(books, output)

}

func init() {
	getCommand.Flags().String(getBooksOutputFlag, "", "Set data output. Prints to console if empty.")
}

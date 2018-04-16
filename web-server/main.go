package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "web",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("from").Value.String()
		serveSite(path)
	},
}

var helpCommand = &cobra.Command{
	Use: "help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This utility runs static web-site on localhost:80 from folder, described in --from flag")
	},
}

func serveSite(dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	log.Println(http.Dir(dir))

	log.Println("Running site")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
	}
}

func init() {
	rootCommand.AddCommand(helpCommand)
	rootCommand.Flags().String(
		"from",
		"static",
		"Set folder with static content",
	)
}

func main() {
	// log.Println(os.Args[1])
	// serveSite(os.Args[1])
	rootCommand.Execute()
}

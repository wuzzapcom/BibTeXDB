package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"wuzzapcom.io/Coursework/pkg/bibtex"

	"wuzzapcom.io/Coursework/pkg/fetchers"
)

//CLI ..
type CLI struct {
	GoogleAPIToken string
	RunAPI         bool
	apiConfig      APIConfiguration
	commands       CommandsCLI
}

//APIConfiguration ..
type APIConfiguration struct {
}

//CommandsCLI ..
type CommandsCLI struct {
	generate string
	search   string
}

//ParseCLIArguments ..
func (cli *CLI) ParseCLIArguments() {
	token := flag.String("token", "", "Token for Google Books api. Exit if empty")

	runAPI := flag.Bool("runAPI", false, "Run JSON api, false by default")

	generateReport := flag.String("generateReport", "", "Used for generation of reports for subject in format \"SubjectName:YYYY\"")

	search := flag.String("searchFor", "", "Search and print result in BibTeX format. Enter in format \"Title; Author\" or you can pick just one")

	flag.Parse()

	if *token == "" {
		cli.exitWithMessage("Provided no Google Books token")
	}

	cli.GoogleAPIToken = *token
	cli.RunAPI = *runAPI

	if !cli.RunAPI {
		cli.commands = CommandsCLI{
			generate: *generateReport,
			search:   *search,
		}
	}

}

//ApplyCommand ..
func (cli *CLI) ApplyCommand() {

	if cli.commands.search != "" {
		cli.search()
	}

}

func (cli *CLI) search() {

	requestedInfo := strings.Split(cli.commands.search, "; ")

	fetcher := fetchers.GoogleFetcher{
		APIToken: cli.GoogleAPIToken,
	}

	var result bibtex.Items
	var err error
	if len(requestedInfo) == 1 {
		result, err = fetcher.FetchWithString(requestedInfo[0])
	} else if len(requestedInfo) == 2 {
		result, err = fetcher.FetchWithTitleAndAuthor(requestedInfo[0], requestedInfo[1])
	} else {
		cli.exitWithMessage("Failed parsing of search query")
	}

	if err != nil {
		cli.exitWithMessage(err.Error())
	}

	fmt.Println(result)

}

//ExitWithMessage ..
func (cli CLI) exitWithMessage(message string) {
	fmt.Println(message)
	os.Exit(-1)
}

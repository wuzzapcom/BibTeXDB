package main

import (
	// "flag"
	"fmt"

	"wuzzapcom.io/Coursework/pkg/cli"
	"wuzzapcom.io/Coursework/pkg/fetchers"
)

func main() {

	cliParser := cli.CLI{}

	cliParser.ParseCLIArguments()

	cliParser.ApplyCommand()

}

func testFetch(googleAPIToken string) {

	fetcher := fetchers.GoogleFetcher{
		APIToken: googleAPIToken,
	}
	result, err := fetcher.FetchWithString("канатников")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

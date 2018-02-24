package main

import (
	"flag"
	"fmt"

	"wuzzapcom.io/Coursework/pkg/database"

	"wuzzapcom.io/Coursework/pkg/fetchers"
)

func main() {

	googleAPIToken := flag.String("googleAPIToken", "", "API token for search in Google Books.")

	flag.Parse()

	fetcher := fetchers.GoogleFetcher{
		APIToken: *googleAPIToken,
	}
	result, err := fetcher.FetchWithString("канатников")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	database.TestConn()
}

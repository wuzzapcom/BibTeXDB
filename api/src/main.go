package main

import (
	"flag"
	"fmt"
	"wuzzapcom/Coursework/api/src/database"
	"wuzzapcom/Coursework/api/src/fetchers"
	"wuzzapcom/Coursework/api/src/restful"
)

func main() {
	googleToken := flag.String("googleToken", "", "Token for Google Books API.")
	port := flag.Int("postgrePort", 32770, "Set port for PostgreSQL connection")
	flag.Parse()
	if *googleToken == "" {
		fmt.Println("No Google Books API token")
		return
	}

	database.Configuration.Port = *port

	fetcher := fetchers.GoogleFetcher{
		APIToken: *googleToken,
	}

	restful.Run(fetcher)
}

func testFetch(googleAPIToken string) {

	fetcher := fetchers.GoogleFetcher{
		APIToken: googleAPIToken,
	}
	result, err := fetcher.FetchWithString("канатников")
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		return
	}
	fmt.Println(result)

}

package main

import (
	"flag"
	"fmt"
	"wuzzapcom/Coursework/api/src/fetchers"
	"wuzzapcom/Coursework/api/src/restful"
)

func main() {
	googleToken := flag.String("googleToken", "", "Token for Google Books API.")
	flag.Parse()
	if *googleToken == "" {
		fmt.Println("No Google Books API token")
		return
	}

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

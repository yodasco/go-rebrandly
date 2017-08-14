package main

import (
	"fmt"
	"os"

	"github.com/yodasco/rebrandly"
)

func main() {
	key := os.Getenv("REBRANDLY_KEY")

	request, err := rebrandly.InitListLinks(
		// Filters:
		false, // not favorite
		"",    // any status
		"",    // all domains
		rebrandly.OrderPagination{}, // defaul ordering and pagination
	)
	if err != nil {
		panic(err)
	}

	answer, err := request.SendRequest(key)
	if err != nil {
		panic(err)
	}

	list := answer.(rebrandly.LinkRequestList)
	fmt.Println("Links")
	fmt.Println("=====")
	for _, link := range list {
		fmt.Println("\tLink Details")
		fmt.Println("\t------------")
		fmt.Println("\tID -", link.ID)
		fmt.Println("\tStatus -", link.Status)
		fmt.Println("\tTitle -", link.Title)
		fmt.Println("\tSlagTag -", link.SlashTag)
		fmt.Println("\tDestination -", link.Destination)
		fmt.Println("\tShortURL -", link.ShortURL)
		fmt.Println("\tClicks -", link.Clicks)
		fmt.Println("\tLastClickAt -", link.LastClickAt)
		fmt.Println("\tFavourite -", link.Favourite)
		fmt.Println("\tForwardParameters -", link.ForwardParameters)
		fmt.Println("")
	}
}

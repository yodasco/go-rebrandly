package main

import (
	"fmt"
	"os"

	"github.com/yodasco/rebrandly"
)

func main() {
	key := os.Getenv("REBRANDLY_KEY")
	linkID := os.Getenv("LINK_ID")

	request, err := rebrandly.InitLinkDetails(linkID)
	if err != nil {
		panic(err)
	}

	answer, err := request.SendRequest(key)
	if err != nil {
		panic(err)
	}

	link := answer.(rebrandly.LinkRequest)

	fmt.Println("Link Details")
	fmt.Println("============")
	fmt.Println("ID -", link.ID)
	fmt.Println("Status -", link.Status)
	fmt.Println("Title -", link.Title)
	fmt.Println("SlagTag -", link.SlashTag)
	fmt.Println("Destination -", link.Destination)
	fmt.Println("ShortURL -", link.ShortURL)
	fmt.Println("Clicks -", link.Clicks)
	fmt.Println("LastClickAt -", link.LastClickAt)
	fmt.Println("Favourite -", link.Favourite)
	fmt.Println("ForwardParameters -", link.ForwardParameters)
}

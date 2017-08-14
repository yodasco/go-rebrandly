package main

import (
	"fmt"
	"os"

	"github.com/yodasco/rebrandly"
)

func main() {
	key := os.Getenv("REBRANDLY_KEY")

	request, err := rebrandly.InitCreateLink(
		// destination
		"https://www.youtube.com/watch?v=x53JHab2ng8",
		// do not create our own slagTag
		"")

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
}

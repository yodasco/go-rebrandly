package main

import (
	"fmt"
	"os"

	"github.com/yodasco/rebrandly"
)

func main() {
	key := os.Getenv("REBRANDLY_KEY")
	domainID := os.Getenv("REBRANDLY_DOMAIN_ID")

	request, err := rebrandly.InitCreateLinkEx(rebrandly.LinkRequest{
		Destination: "https://www.youtube.com/watch?v=x53JHab2ng8",
		Title:       "Cute Gophers",
		// Use custom domain, rather then rebrand.ly
		Domain: rebrandly.DomainRequest{
			ID:  domainID,
			Ref: fmt.Sprintf("domains/%s", domainID),
		},
	})

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

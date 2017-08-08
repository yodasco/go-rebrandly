package rebrandly

import "time"

// The following section holds the structure for All the requests and responces
// that defined at the model section by Rebrandly documentation located at:
// https://developers.rebrandly.com/docs/model-overview

// The default content type to use
const contentType = "application/json"

// DomainTypes holds an "enum" of allowed types
type DomainTypes string

// Enumeration values for DomainTypes
const (
	DomainTypeService DomainTypes = "service"
	DomainTypeUser    DomainTypes = "user"
)

// DomainRequest holds the main domain fields for a request
// JSON example for such request
// {
//   "id": "xxxxxxxxxxxxxxxxx",
//   "fullName": "brand.cool",
//   "topLevelDomain": "cool",
//   "createdAt": "2016-07-01T13:12:22.000Z",
//   "updatedAt": "2016-07-03T13:17:50.000Z",
//   "type": "user",
//   "active": false
// }
type DomainRequest struct {
	// Unique identifier for the branded domain
	ID string `json:"id"`
	// Full name of the branded domain
	FullName string `json:"fullName"`
	// The top level domain part of the branded domain name
	TopLevelDomain string `json:"topLevelDomain"`
	// UTC creation date/time of the branded domain
	CreatedAt time.Time `json:"createdAt"`
	// UTC last update date/time of the branded domain
	UpdatedAt time.Time `json:"updatedAt"`
	// Branded domain type
	Type DomainTypes `json:"type"`
	// Whether the branded domain can be used or not to create branded short links
	Active bool `json:"active"`
}

// LinkStatus holds an "enum" of allowed types
type LinkStatus string

// Enumeration values for LinkStatus
const (
	// When a link is alive, it is always active
	LinkStatusActive LinkStatus = "active"
	// A Link is said to be in trashed status when it has been temporarily
	// deleted
	LinkStatusTrashed LinkStatus = "trashed"
)

// LinkRequest holds the main link fields for a request
// {
//   "id": "xxxxxxxxxxxxxxxxx",
//   "title": "The LaFerrari Supercar Convertible Is the New Best Way to Burn $1M | WIRED",
//   "slashtag": "burn10M",
//   "destination": "https://www.wired.com/2016/07/ferrari-laferrari-spider-convertible-photos-specs/",
//   "shortUrl": "rebrand.ly/burn10M",
//   "domain": {
//     "id": "8f104cc5b6ee4a4ba7897b06ac2ddcfb",
//     "fullName": "rebrand.ly"
//   },
//   "status": "active",
//   "createdAt": "2016-07-13T10:54:12.000Z",
//   "updatedAt": "2016-07-13T10:54:12.000Z",
//   "clicks": 42,
//   "lastClickAt": "2016-07-13T10:55:13.000Z",
//   "favourite": false,
//   "forwardParameters": true
// }
type LinkRequest struct {
	// Unique identifier associated with the branded short link
	ID string `json:"id"`
	// A title you assign to the branded short link in order to remember what's
	// behind it
	Title string `json:"title"`
	// The keyword section of your branded short link
	SlashTag string `json:"slashtag"`
	// The destination URL you want your branded short link to point to
	Destination string `json:"destination"`
	// The full branded short link URL, including domain
	ShortURL string `json:"short_url"`
	// A reference to the branded domain's resource of a branded short link
	Domain DomainRequest `json:"domain"`
	// Status of the branded short link.
	Status LinkStatus `json:"status"`
	// The UTC date/time this branded short link was created
	CreatedAt time.Time `json:"createdAt"`
	// The last UTC date/time this branded short link was updated.
	// When created, it matches CreatedAt
	UpdatedAt time.Time `json:"updatedAt"`
	// How many clicks there are on this branded short link so far
	Clicks int64 `json:"clicks"`
	// The UTC date/time this branded short link was last clicked on
	LastClickAt time.Time `json:"lastClickAt"`
	// Whether a link is favourited (loved) or not
	Favourite bool `json:"favourite"`
	// Whether query parameters in short URL will be forwarded to destination URL.
	// E.g. short.link/kw?p=1 with forwardParameters=true will redirect to
	// longurl.com/home/path?p=1, otherwise will redirect to longurl.com/home/path
	// (without query parameters)
	ForwardParameters bool `json:"forwardParameters"`
}

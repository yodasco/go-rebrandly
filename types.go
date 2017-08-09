package rebrandly

import "time"

// The following section holds the structure for All the requests and responces
// that defined at the model section by Rebrandly documentation located at:
// https://developers.rebrandly.com/docs/model-overview

// DomainTypes holds an "enum" of allowed types
type DomainTypes string

// Enumeration values for DomainTypes
const (
	DomainTypeService DomainTypes = "service"
	DomainTypeUser    DomainTypes = "user"
)

// DomainRequest holds the main domain fields for a request
// JSON example for such request
//
//   {
//     "id": "xxxxxxxxxxxxxxxxx",
//     "fullName": "brand.cool",
//     "topLevelDomain": "cool",
//     "createdAt": "2016-07-01T13:12:22.000Z",
//     "updatedAt": "2016-07-03T13:17:50.000Z",
//     "type": "user",
//     "active": false
//   }
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

// DomainRequestList is a list of one or more DomainRequest
type DomainRequestList []DomainRequest

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
// JSON example for such request
//
//   {
//     "id": "xxxxxxxxxxxxxxxxx",
//     "title": "The LaFerrari Supercar Convertible Is the New Best Way to Burn $1M | WIRED",
//     "slashtag": "burn10M",
//     "destination": "https://www.wired.com/2016/07/ferrari-laferrari-spider-convertible-photos-specs/",
//     "shortUrl": "rebrand.ly/burn10M",
//     "domain": {
//       "id": "8f104cc5b6ee4a4ba7897b06ac2ddcfb",
//       "fullName": "rebrand.ly"
//     },
//     "status": "active",
//     "createdAt": "2016-07-13T10:54:12.000Z",
//     "updatedAt": "2016-07-13T10:54:12.000Z",
//     "clicks": 42,
//     "lastClickAt": "2016-07-13T10:55:13.000Z",
//     "favourite": false,
//     "forwardParameters": true
//   }
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

// LinkRequestList holds a list of LinkRequest
type LinkRequestList []LinkRequest

// AccountLimit holds the structure for limits at the main Account structure
type AccountLimit struct {
	// How many resources of the given type used
	Used int64 `json:"used"`
	// How many resources of the given type the account is allowing
	Max int64 `json:"max"`
}

// AccountLimitName holds an "enum" of limitation names
type AccountLimitName string

// Holds the enum values for AccountLimitName
const (
	AccountLimitNameLinks     AccountLimitName = "links"
	AccountLimitNameDomains   AccountLimitName = "domains"
	AccountLimitNameTeamMates AccountLimitName = "teammates"
	AccountLimitNameTags      AccountLimitName = "tags"
	AccountLimitNameScripts   AccountLimitName = "scripts"
)

// AccountSubscription holds subscription information on specific category
type AccountSubscription struct {
	// Category the account's plan belongs to
	Category string `json:"category"`
	// UTC subscription date/time of the account's current plan
	CreatedAt time.Time `json:"createdAt"`
	// UTC expiration date/time of the account's current plan, when plan's
	// category is not free
	ExpiredAt time.Time `json:"expiredAt"`
	// Account's resources usage and limits: how many links/domains/tags/etc
	// created so far and which are the maximum limits
	Limits map[AccountLimitName]AccountLimit `json:"limits"`
}

// AccountRequest holds the fields for account request
// JSON example for such request
//
//   {
//     "id": "xxxxxxxxxxxxxxxxx",
//     "fullName": "Stanford University",
//     "username": "fake@stanford.edu",
//     "email": "fake@stanford.edu",
//     "avatarUrl": "https://d3e7f5z1blhqw4.cloudfront.net/avatars/364381e1-963e-460a-9a6b-a16e86d196a2",
//     "createdAt": "2016-07-13T10:54:12.000Z",
//     "subscription": {
//       "category": "free",
//       "createdAt": "2016-07-13T10:54:12.000Z",
//       "limits": {
//         "links": {
//           "used": 7504,
//           "max": 10000
//         },
//         "domains": {
//           "used": 17,
//           "max": 100
//         },
//         "teammates": {
//           "used": 5,
//           "max": 100
//         },
//         "tags": {
//           "used": 0,
//           "max": 50
//         },
//         "scripts": {
//           "used": 0,
//           "max": 50
//         }
//       }
//     }
//   }
type AccountRequest struct {
	// Unique identifier of the account
	ID string `json:"id"`
	// Username used in login
	Username string `json:"username"`
	// Contact email of the account
	Email string `json:"email"`
	// Full name of the account owner
	FullName string `json:"fullName"`
	// URL of the account avatar
	AvatarURL string `json:"avatarUrl"`
	// UTC creation date/time of the account
	CreatedAt time.Time `json:"createdAt"`
	// Set of feature/limits info related to the account and its plan
	Subscription AccountSubscription `json:"subscription"`
}

// ResourceRequest is a means to connect two resources together
// JSON example for such request
//
//   {
//     "id": "xxxx12433xxx135555",
//     "ref": "/domains/8f104cxxxxxxxa7897bxxxxxxxfb"
//   }
type ResourceRequest struct {
	// Unique identifier of the original resource
	ID string `json:"id"`
	// API path to resource details
	Ref string `json:"ref"`
}

// CountRequest is a request for counters
//
// JSON example:
//   {
//      "count": 42
//   }
type CountRequest struct {
	Count int64 `json:"count"`
}

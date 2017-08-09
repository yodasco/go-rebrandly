package rebrandly

// InitCreateLinkEx initialize the Request struct with parameters for creating
// a link.
// The function uses LinkRequest struct to better control the creation of a new
// link.
func InitCreateLinkEx(fields LinkRequest) (Request, error) {
	err := validDestination(fields.Destination)
	if err != nil {
		return Request{}, err
	}

	return Request{}, nil
}

// InitCreateLink initialize the Request struct with parameters for creating
// a link.
// The initialization is only with mandatory fileds.
// For advanced initialization, use the InitCreateLinkEx func instead
func InitCreateLink(destination string) (Request, error) {
	return InitCreateLinkEx(LinkRequest{
		Destination: destination,
	})
}

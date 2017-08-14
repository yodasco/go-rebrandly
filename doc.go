/*
Package rebrandly General Overview

The package implements most of the APIs by rebrandly.com.
The implementation provides full control over links by providing CRUD operations, listing them, and counting them.

The library contacts rebrandly by generating a `Request` struct which in turn provides a connection and body information regarding the connection.

There are several helper functions to initialize the `Request` struct by
adding the proper configuration inside the `Request` struct.

When an operation of the API needs to send a JSON body, there are two helper
functions implementation:
  * `InitXxxEx`
  * `InitXxx`

The Ex suffix stand for extended, and provides a means to fully control the
JSON that is going to be sent out.

The `InitXxx` function, requires only the minimal parameters that are usually
mandatory by the API in order to function, but does not add extra fields,
focusing only on the task rather then the extra functionality.


Basic Usage
-----------

Here is the most simple means to create a new link.

    link, err := rebrandly.InitListLinks(
       "https://www.youtube.com/watch?v=x53JHab2ng8", "sdd12Wa")
    if err != nil {
       panic(err)
    }

    details, err := link.SendRequest("1234567890")
	if err != nil {
		panic(err)
	}

	fmt.Println("details.ShortURL: ", details.(rebrandly.LinkRequest).ShortURL)


If everything went well, `details` is now a `LinkRequest` struct that
holds information regarding the link of https://rebrand.ly/sdd12Wa
with full details about it.

*/
package rebrandly

Rebrandly API
=============

An implementation in Go (golang) for the [REST API](https://developers.rebrandly.com/docs/) of Rebrandly.


Package rebrandly General Overview
===================================

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


Error handling
----------------

Functions in this package typically yield two types of errors:
   1. Programmatic errors, using Go's `error` type
   2. REST errors - Errors that come up as a result of using the API itself

The 2nd type of errors, are typically structs parsed from the JSON returned by rebrandly.
They are listed in the file `error_types.go` and are documented there.

// TODO(ido): I'm not sure it's a wise design decision, let's branstorm about that. 
Therefore when using the API it is important to realize that even when the API returns a non-error result, the client of the API must look into the returned struct to see if it encapsulates a REST error returned from rebrandly. 

As convinience, the function `IsErrorStruct` takes such struct and determines if it is a REST error or not.

Basic Usage
-----------

Here is the most simple means to create a new link.

    link, err := rebrandly.InitListLinks(
       "https://www.youtube.com/watch?v=x53JHab2ng8", "sdd12Wa")
    if err != nil {
       panic(err)
    }
  
    details, err := link.SendRequest("1234567890")
    
    //TODO(ido): Shouldn't you use the IsErrorStruct here?

If everything went well, `details` is now a `LinkRequest` struct that
holds information regarding the link of https://rebrand.ly/sdd12Wa
with full details about it.

If there was an error returned by the server, then an error struct will be
returned.

Any other type of error will be placed on the `err` variable instead.

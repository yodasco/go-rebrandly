Rebrandly API
=============

An implementation in Go (golang) for the [REST API](https://developers.rebrandly.com/docs/) of Rebrandly.


Package rebrandly General Overview
===================================

The package implement most of the API by rebrandly.com.
The implementation provide full control over links, by doing CRUD operations,
listing them, and counting them.

The library contact rebrandly by generating a `Request` struct tat provides
connection and body information regarding the connection.

There are several helper functions to initialize the `Request` struct by
adding the proper configuration inside the `Request` struct.

When an operation of the API need to send JSON body, there are two helper
functions implementation:
  * InitxxxxEx
  * Initxxxx

The Ex suffix stand for extended, and provide a mean to fully control the
JSON that is going to be sent.

The counter part function, takes only minimal parameters that are usually
mandatory by the API in order to function, but does not add extra fields,
focusing only on the task rather then the extra functionality.


How errors works
----------------

The package control two types of errors:
   1. Problematically errors, using Go's error type
   2. REST errors - Errors arising from using the API itself

The 2nd error types, are struct based, that parsed from JSON return.
They are located under the `error_types.go` file, and documented to explain
what they represent

Before explaining how to use the library, it is important to understand that
the struct that is returned can be the needed JSON content, or an error
content.

There is an helper function named IsErrorStruct that takes a struct and check
to see if it is an error struct, to be able to identify it easily.


Basic Usage
-----------

Here is the most simple means to create a new link.

    link, err := rebrandl.yInitListLinks(
       "https://www.youtube.com/watch?v=x53JHab2ng8", "sdd12Wa")
    if err != nil {
       panic(err)
    }
  
    details, err := link.SendRequest("1234567890")

If everything went well, `details` is now a LinkRequest struct that
hold information regarding the link of https://rebrand.ly/sdd12Wa
with full details about it.

If there was an error returned by the server, then an error struct will be
returned.

Any other type of error will be placed on the `err` variable instead.



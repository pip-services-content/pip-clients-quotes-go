# Quotes Microservice Client SDK for Golang

This is a Golang client SDK for [pip-services-quotes](https://github.com/pip-services/pip-services-quotes) microservice.
It provides an easy to use abstraction over communication protocols:

* HTTP/REST client

<a name="links"></a> Quick Links:

* [Development Guide](doc/Development.md)
* [API Version 1](doc/GoClientApiV1.md)

## Install

Install runtime and client packages by executing the following command in the command line

```bash
go get github.com/pip-services/pip-services-runtime-go
go get github.com/pip-services/pip-clients-quotes-go
```

## Use

Import packages **github.com/pip-services/pip-services-runtime-go** and **github.com/pip-services/pip-clients-quotes-go/version1**
```go
import (
    ...
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-clients-quotes-go/version1"
)
```

Define client configuration parameters that match configuration of the microservice external API
```go
// Client configuration
config := runtime.NewMapAndSet(
    "transport.type", "http",
    "transport.host", "localhost",
    "transport.port", 8002,
)
```

Instantiate the client and open connection to the microservice
```go
// Create the client instance
client := quotes_v1.NewQuotesRestClient(config)

// Connect to the microservice
err := client.Open()
    
// Work with the microservice
...
```

Now the client is ready to perform operations
```go
// Create a new quote
quote := quotes_v1.Quote{
    Text: map[string]string { "en": "Get in hurry slowly" },
    Author: map[string]string { "en": "Russian proverb" },
    Tags: []string { "time management" },
    Status: "completed",
}

quote, err = client.CreateQuote(quote)
```

```go
// Get the list of quotes on 'time management' topic
quotePage := client.GetQuotes(
    runtime.NewFilterParamsAndSet(
        "tags", "time management",
        "status", "completed",
    ),
    runtime.NewPagingParams(0, 10),
)
```    

## Acknowledgements

This client SDK was created and currently maintained by *Sergey Seroukhov*.


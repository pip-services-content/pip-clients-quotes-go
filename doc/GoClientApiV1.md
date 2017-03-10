# Client API (version 1) <br/> Quotes Microservices Client SDK for Golang

Go client API for Quotes microservice is a thin layer on the top of
communication protocols. It hides details related to specific protocol implementation
and provides high-level API to access the microservice for simple and productive development.

* [Installation](#install)
* [Getting started](#get_started)
* [MultiString map](#class1)
* [Quote struct](#class2)
* [QuotePage struct](#class3)
* [IQuoteClient interface](#interface)
    - [Init()](#operation1)
    - [Open()](#operation2)
    - [Close()](#operation3)
    - [GetQuotes()](#operation4)
    - [GetRandomQuote()](#operation5)
    - [GetQuoteById()](#operation6)
    - [CreateQuote()](#operation7)
    - [UpdateQuote()](#operation8)
    - [DeleteQuote()](#operation9)
* [QuotesRestClient struct](#client_rest)

## <a name="install"></a> Installation

To work with the client SDK import **github.com/pip-services/pip-clients-quotes/version1** package

```bash
go get github.com/pip-services/pip-clients-quotes/version1
```

Then go to the package folder and install all secondary dependencies

```bash
go get ./...
```

## <a name="get_started"></a> Getting started

This is a simple example on how to work with the microservice using REST client:

```go
package sample

import (
    "fmt"
    "log"
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-clients-quotes-go/version1"
)

func main() {
    // Client configuration
    config := runtime.NewMapAndSet(
        "transport.type", "http",
        "transport.host", "localhost",
        "transport.port", 8002,
    )

    // Create the client instance
    client := quotes_v1.NewQuotesRestClient(config)
    
    // Open client connection to the microservice
    err1 := client.Open()
    if err1 != nil {
        log.Fatalf("Failed to open client: %v", err1)
    }
    fmt.Println("Opened connection");

    // Create a new quote
    quote := quotes_v1.Quote {
        Text: map[string]string { "en": "Get in hurry slowly" },
        Author: map[string]string { "en", "Russian proverb" },
        Tags: []string { "time management" },
        Status: "completed",
    }

    quote, err2 := client.CreateQuote(quote)
    if err2 != nil {
        log.Fatalf("Failed to create quote: %v", err2)
    }
    fmt.Println("Create quote is")
    fmt.Println(quote)

    // Get the list of quotes on 'time management' topic
    quotePage, err3 := client.GetQuotes(
        NewFilterParamsAndSet(
            "tags", "time management",
            "status", "completed",
        ),
        NewPagingParams(0, 10),
    )
    if err3 != nil {
        log.Fatalf("Failedto read quotes: %v", err3)
    }

    fmt.Println("Quotes on time management are");
    fmt.Println(quotesPage.Data)
    
    // Close connection
    client.Close() 
}
```

### <a name="class1"></a> MultiString map[string]string

String map that contains versions in multiple languages

**Properties:**
- en: string - English version of the string
- sp: string - Spanish version of the string
- de: string - German version of the string
- fr: string - Franch version of the string
- pt: string - Portuguese version of the string
- ru: string - Russian version of the string
- .. - other languages can be added here

### <a name="class2"></a> Quote struct

Represents an inspirational quote

**Properties:**
- ID: string - unique quote id
- Text: MultiString - quote text in different languages
- Author: MultiString - name of the quote author in different languages
- Status: string - editing status of the quote: 'new', 'writing', 'translating', 'completed' (default: 'new')
- Tags: string[] - (optional) search tags that represent topics associated with the quote
- AllTags: string[] - (read only) explicit and hash tags in normalized format for searching  

### <a name="class3"></a> QuotePage struct

Represents a paged result with subset of requested quotes

**Properties:**
- Data: []*Quote - array of retrieved Quote references
- Total: Integer - total number of objects in retrieved resultset

## <a name="interface"></a> IQuotesClient interface

IQuotesClient as a common interface across all client implementations. 

```go
type IQuotesClient interface {
    Init(refs *References) error
    Open() error
    Close() error
    GetQuotes(filter *FilterParams, paging *PagingParams) (*QuotesDataPage, error)
    GetRandomQuote(filter *FilterParams) (*Quote, error)
    GetQuoteById(quoteID string) (*Quote, error)
    CreateQuote(quote *Quote) (*Quote, error)
    UpdateQuote(quoteID string, quote *DynamicMap) (*Quote, error)
    DeleteQuote(quoteID string) error
}
```

### <a name="operation1"></a> Init(refs)

Initializes client references. This method is optional. It is used to set references 
to logger or performance counters.

**Arguments:**
- refs: *References - references to other components 
  - Log: ILog - reference to logger
  - Counters: ICounters - reference to performance counters

**Returns**
  - error - Occured error

### <a name="operation2"></a> Open()

Opens connection to the microservice

**Returns**
  - error - Occured error

### <a name="operation3"></a> Close()

Closes connection to the microservice

**Returns**
  - error - Occured error

### <a name="operation4"></a> GetQuotes(filter, paging)

Retrieves a collection of quotes according to specified criteria

**Arguments:** 
- filter: *FilterParams - filter parameters
  - tags: string[] - (optional) list tags with topic names
  - status: string - (optional) quote editing status
  - author: string - (optional) author name in any language 
  - except_ids: string[] - (optional) quote ids to exclude 
- paging: *PagingParams - paging parameters
  - skip: int - (optional) start of page (default: 0). Operation returns paged result
  - take: int - (optional) page length (max: 100). Operation returns paged result
  - paging: bool - (optional) true to enable paging and return total count

**Returns**
  - *QuotesDataPage - retrieved quotes in paged format
  - error - Occured error

### <a name="operation5"></a> GetRandomQuote(filter)

Retrieves a random quote from filtered resultset

**Arguments:** 
- filter: *FilterParams - filter parameters
  - tags: string[] - (optional) list tags with topic names
  - status: string - (optional) quote editing status
  - author: string - (optional) author name in any language
  - except_ids: string[] - (optional) quote ids to exclude
  
**Returns** 
  - *Quote - random quote, null if object wasn't found 
  - error - Occured error

### <a name="operation6"></a> GetQuoteById(quoteID)

Retrieves a single quote specified by its unique id

**Arguments:** 
- quoteID: string - unique Quote id

**Returns**
- *Quote - retrieved quote, null if object wasn't found 
- error - Occured error

### <a name="operation7"></a> CreateQuote(quote)

Creates a new quote

**Arguments:** 
- quote: *Quote - Quote object to be created. If object id is not defined it is assigned automatically.

**Returns**
- *Quote - created quote object
- error - Occured error

### <a name="operation8"></a> UpdateQuote(quoteID, quote)

Updates quote specified by its unique id

**Arguments:** 
- quoteID: string - unique quote id
- quote: DynamicMap - quote object with new values. Partial updates are supported

**Returns**
- *Quote - updated quote object 
- error - Occured error

### <a name="operation9"></a> DeleteQuote(quoteID)

Deletes quote specified by its unique id

**Arguments:** 
- quoteID: string - unique quote id

**Returns**
- error - Occured error

## <a name="client_rest"></a> QuotesRestClient class

QuotesRestClient is a client that implements HTTP/REST protocol

```go
type QuotesRestClient struct {
    RestClient
    ...
}

func NewQuotesRestClient(config *DynamicMap) *QuotesRestClient {
    ...
}
```

**Constructor config properties:** 
- transport: Object - HTTP transport configuration options
  - type: String - HTTP protocol - 'http' or 'https' (default is 'http')
  - host: String - IP address/hostname binding (default is '0.0.0.0')
  - port: int - HTTP port number

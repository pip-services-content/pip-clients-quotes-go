package quotes_v1

import (
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/deps"            
)

type QuotesRestClient struct {
    deps.RestClient
}

func NewQuotesRestClient(config *runtime.DynamicMap) *QuotesRestClient {
    return &QuotesRestClient { RestClient: *deps.NewRestClient("Quotes.RestClient", config) }
}

func (c *QuotesRestClient) GetQuotes(filter *runtime.FilterParams, paging *runtime.PagingParams) (*QuotesDataPage, error) {
    timing := c.Instrument("Quotes.GetQuotes")
    defer func() { timing.EndTiming() }()
    
    outputPage := NewEmptyQuotesDataPage()
    _, err := c.Call("GET", "/quotes" + c.GetPagingAndFilterParams(filter, paging), nil, outputPage)
    
    if err != nil { return nil, err }
           
    return outputPage, nil
}

func (c *QuotesRestClient) GetRandomQuote(filter *runtime.FilterParams) (*Quote, error) {
    timing := c.Instrument("Quotes.GetRandomQuote")
    defer func() { timing.EndTiming() }()
    
    outputQuote := NewEmptyQuote()
    result, err := c.Call("GET", "/quotes/random" + c.GetPagingAndFilterParams(filter, nil), nil, outputQuote)
    
    if err != nil { return nil, err }
    if result == nil { return nil, nil }
           
    return outputQuote, nil
}

func (c *QuotesRestClient) GetQuoteById(quoteID string) (*Quote, error) {
    timing := c.Instrument("Quotes.GetQuoteById")
    defer func() { timing.EndTiming() }()
    
    outputQuote := NewEmptyQuote()
    result, err := c.Call("GET", "/quotes/" + quoteID, nil, outputQuote)
    
    if err != nil { return nil, err }
    if result == nil { return nil, nil }
           
    return outputQuote, nil
}

func (c *QuotesRestClient) CreateQuote(quote *Quote) (*Quote, error) {
    timing := c.Instrument("Quotes.CreateQuote")
    defer func() { timing.EndTiming() }()

    outputQuote := NewEmptyQuote()
    _, err := c.Call("POST", "/quotes", quote, outputQuote)
    
    if err != nil { return nil, err }
       
    return outputQuote, nil
}

func (c *QuotesRestClient) UpdateQuote(quoteID string, quote *runtime.DynamicMap) (*Quote, error) {
    timing := c.Instrument("Quotes.UpdateQuote")
    defer func() { timing.EndTiming() }()
    
    outputQuote := NewEmptyQuote()
    result, err := c.Call("PUT", "/quotes/" + quoteID, quote, outputQuote)
    
    if err != nil { return nil, err }
    if result == nil { return nil, nil }
           
    return outputQuote, nil
}

func (c *QuotesRestClient) DeleteQuote(quoteID string) error {
    timing := c.Instrument("Quotes.DeleteQuote")
    defer func() { timing.EndTiming() }()
    
    _, err := c.Call("DELETE", "/quotes/" + quoteID, nil, nil)
               
    return err
}

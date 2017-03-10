package deps

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-clients-quotes-go/version1"
)

type QuotesClientFixture struct {
    client quotes_v1.IQuotesClient
}

func NewQuotesClientFixture(client quotes_v1.IQuotesClient) *QuotesClientFixture {
    return &QuotesClientFixture { client: client }    
}

var QUOTE1 *quotes_v1.Quote = createQuote("", "Text 1", "Author 1")
var QUOTE2 *quotes_v1.Quote = createQuote("", "Text 2", "Author 2")

func createQuote(id, text, author string) *quotes_v1.Quote {
    return &quotes_v1.Quote {
        ID: "",
        Text: map[string]string { "en": text },
        Author: map[string]string { "en": author },
        Status: "new",
        Tags: []string {},
        AllTags: []string {},
    }
}

func (c *QuotesClientFixture) TestCrudOperations(t *testing.T) {
    // Create one quote
    quote1, err := c.client.CreateQuote(QUOTE1)    
    assert.Nil(t, err)
    
    assert.NotNil(t, quote1)
    assert.NotEmpty(t, quote1.ID)
    assert.Equal(t, QUOTE1.Text["en"], quote1.Text["en"])
    assert.Equal(t, QUOTE1.Author["en"], quote1.Author["en"])

    // Create another quote
    quote2, err := c.client.CreateQuote(QUOTE2)    
    assert.Nil(t, err)
    
    assert.NotNil(t, quote2)
    assert.NotEmpty(t, quote2.ID)
    assert.Equal(t, QUOTE2.Text["en"], quote2.Text["en"])
    assert.Equal(t, QUOTE2.Author["en"], quote2.Author["en"])
    
    // Get all quotes
    quotes, err2 := c.client.GetQuotes(nil, nil)
    assert.Nil(t, err2)
    assert.NotNil(t, quotes)
    assert.NotNil(t, quotes.Data)
    assert.True(t, len(quotes.Data) >= 2)
        
    // Update the quote
    quoteData := runtime.NewMapAndSet("text.en", "Updated Text 1")
    quote1, err = c.client.UpdateQuote(quote1.ID, quoteData)
    assert.Nil(t, err)
    assert.NotNil(t, quote1)
    assert.Equal(t, "Updated Text 1", quote1.Text["en"])

    // Delete the quote #1
    err = c.client.DeleteQuote(quote1.ID)
    assert.Nil(t, err)

    // Delete the quote #2
    err = c.client.DeleteQuote(quote2.ID)
    assert.Nil(t, err)
    
    // Try to get deleted quote
    quote1, err = c.client.GetQuoteById(quote1.ID)
    assert.Nil(t, err)
    assert.Nil(t, quote1) 
}
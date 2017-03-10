package quotes_v1

import (
    "github.com/pip-services/pip-services-runtime-go"
)

type IQuotesClient interface {
    GetQuotes(filter *runtime.FilterParams, paging *runtime.PagingParams) (*QuotesDataPage, error)
    GetRandomQuote(filter *runtime.FilterParams) (*Quote, error)
    GetQuoteById(quoteID string) (*Quote, error)
    CreateQuote(quote *Quote) (*Quote, error)
    UpdateQuote(quoteID string, quote *runtime.DynamicMap) (*Quote, error)
    DeleteQuote(quoteID string) error
}
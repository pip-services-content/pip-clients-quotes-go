package quotes_v1

type QuotesDataPage struct {
    Total *int         `json:"total"`
    Data []*Quote       `json:"data"`
}

func NewEmptyQuotesDataPage() *QuotesDataPage {
    return &QuotesDataPage{}
}
func NewQuotesDataPage(total *int, data []*Quote) *QuotesDataPage {
    return &QuotesDataPage{ Total: total, Data: data }
}
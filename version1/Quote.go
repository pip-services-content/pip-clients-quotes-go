package quotes_v1

type Quote struct {
    ID      string  `json:"id" bson:"_id"`
    Text    map[string]string  `json:"text" bson:"text"`
    Author  map[string]string  `json:"author" bson:"author"`
    Status  string `json:"status" bson:"status"`
    Tags    []string `json:"tags" bson:"tags"`
    AllTags []string `json:"all_tags" bson:"all_tags"`
}

func NewEmptyQuote() *Quote {
    return &Quote {}
}

func NewQuote(id string, text map[string]string, author map[string]string, status string, tags []string, allTags []string) *Quote {
    return &Quote { ID: id, Text: text, Author: author, Status: status, Tags: tags, AllTags: allTags }
}

func (c *Quote) GetID() string {
    return c.ID
}

func (c *Quote) SetID(id string) {
    c.ID = id
}
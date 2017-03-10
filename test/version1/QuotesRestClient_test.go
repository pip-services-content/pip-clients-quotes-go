package deps

import (    
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/build"    
    "github.com/pip-services/pip-clients-quotes-go/version1"    
)

type QuotesRestClientTest struct {
    suite.Suite
    
    client *quotes_v1.QuotesRestClient
    refs *runtime.References
    fixture *QuotesClientFixture
}

func (suite *QuotesRestClientTest) SetupSuite() {
    restConfig := runtime.NewMapAndSet(
        "type", "rest",
        "transport.type", "http",
        "transport.host", "localhost",
        "transport.port", 8002,
    )
    suite.client = quotes_v1.NewQuotesRestClient(restConfig)
    
    suite.fixture = NewQuotesClientFixture(suite.client)
    
    suite.refs = runtime.NewReferences().WithDeps("quotes", suite.client)
           
    err := build.LifeCycleManager.InitAndOpen(suite.refs)
    assert.Nil(suite.T(), err)
}

func (suite *QuotesRestClientTest) TearDownSuite() {
    err := build.LifeCycleManager.Close(suite.refs)
    assert.Nil(suite.T(), err)
}

func (suite *QuotesRestClientTest) TestCrudOperations() {
    suite.fixture.TestCrudOperations(suite.T())
}

func TestQuotesRestClientTestSuite(t *testing.T) {
    suite.Run(t, new(QuotesRestClientTest))
}
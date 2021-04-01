package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/benweissmann/memongo"
	"github.com/futuregerald/futureauth-go/src/functions/futureauth"
	"github.com/kamva/mgm/v3"

	"github.com/stretchr/testify/assert"
)

type mockDBClient struct {
}

// TODO add some basic model validationt to mock save method and return error if fails
func (c *mockDBClient) Save(m mgm.Model) error {
	return nil
}

func TestGetSuccess(t *testing.T) {
	body := `{
    "email": "test.account@testing.com",
	"password":"testingsdf",
	"userMetadata": {"random":"data"},
	"appMetadata": {"random":"data"},
	"roles": ["admin","user"]
	}`
	req, err := http.NewRequest("POST", "/", strings.NewReader(body))
	assert.NoError(t, err)
	mongoServer, err := memongo.Start("4.0.5")
	assert.NoError(t, err)
	defer mongoServer.Stop()
	err = futureauth.New(mongoServer.URI())
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LambdaHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

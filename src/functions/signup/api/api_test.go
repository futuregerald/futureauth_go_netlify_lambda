package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

	rr := httptest.NewRecorder()
	dbClient := &mockDBClient{}
	a := New(dbClient)
	handler := http.HandlerFunc(a.LambdaHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

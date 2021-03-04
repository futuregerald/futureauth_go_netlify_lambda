package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	client := New()
	handler := http.HandlerFunc(client.LambdaHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

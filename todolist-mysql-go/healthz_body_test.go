package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthz_Body(t *testing.T) {
	request, _ := http.NewRequest("GET", "/healthz", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	assert.Equal(t, `{"alive": true}`, response.Body.String(), "Incorrect body found")
}

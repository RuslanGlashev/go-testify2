package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equal(t, status, http.StatusOK)
	require.NotEmpty(t, responseRecorder.Body)
}
func testLenCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req) // здесь нужно создать запрос к сервису
	body := responseRecorder.Body.String()
	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, body, "wrong city value")
}

func testCount(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req) // здесь нужно создать запрос к сервису
	body := responseRecorder.Body.String()

	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)
}

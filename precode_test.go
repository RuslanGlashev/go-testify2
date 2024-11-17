package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder)

	totalCount := 2

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenOtherCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=astana", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expectedError := "wrong city value"
	body := responseRecorder.Body.String()
	assert.Contains(t, body, expectedError)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)

}

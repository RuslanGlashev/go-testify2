package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=murmansk", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expected := "wrong city value"

	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	totalCount := len(cafeList["moscow"])

	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	listOfCafe := strings.Split(responseRecorder.Body.String(), ",")

	assert.Len(t, listOfCafe, totalCount)
}

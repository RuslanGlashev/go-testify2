package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMainHandlerStatusOk()
// Проверка корректности запроса
func TestMainHandlerStatusOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, http.StatusOK, status)

	body := responseRecorder.Body
	assert.NotEmpty(t, body)
}

// TestMainHandlerWrongCityValue()
// Тест на наличие города
func TestMainHandlerWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=krasnodar", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	message := "wrong city value"
	assert.Equal(t, message, responseRecorder.Body.String())
}

// TestMainHandlerWhenCountMoreThanTotal()
// Тест на count > totalCount
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, list, totalCount)

}

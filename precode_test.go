package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/?city=moscow&count="+strconv.Itoa(totalCount+1), nil)
	require.NoError(t, err, "Не удалось создать запрос")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Неверный код ответа")

	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Неверное тело ответа")
}

func TestMainHandlerValidRequest(t *testing.T) {
	// Корректный запрос с count=2 и городом moscow
	req, err := http.NewRequest("GET", "/?city=moscow&count=2", nil)
	require.NoError(t, err, "Не удалось создать запрос")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Неверный код ответа")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа пустое")

	expectedBody := "Мир кофе,Сладкоежка"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Неверное тело ответа")
}

func TestMainHandlerWrongCity(t *testing.T) {
	// Некорректный запрос с неподдерживаемым городом
	req, err := http.NewRequest("GET", "/?city=unknown&count=2", nil)
	require.NoError(t, err, "Не удалось создать запрос")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Неверный код ответа")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Неверное тело ответа")
}

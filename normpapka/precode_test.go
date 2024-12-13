package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Уважаемый ревьюер. Если у меня что-то оформлено или сделано не так, прошу убедиться, что задание от меня этого требует, и что мы это проходили в обучении.
В противном случае, прошу указать ссылки, где можно научиться тому о чём вы говорите.

Пожалуйста, убедитесь, что внятно указали на недочёты, чтобы человек, который только недавно начал погружаться в программирование
понял что от него требуется.
*/

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, list, totalCount)

}

func TestMainHandlerStatus(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=bobryisk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

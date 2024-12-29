package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	// Тест №1: Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	assert.Equal(t, responseRecorder.Code, 200, "Error. Expected code: 200. Actual code: %d", responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String(), "Value should not be empty")

	// Тест №2: Город, который передаётся в параметре `city`, не поддерживается. Сервис возвращает код ответа 400 и ошибку `wrong city value` в теле ответа.
	req = httptest.NewRequest("GET", "/cafe?city=tula&count=5", nil)

	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, req)
	actualAnswer := responseRecorder.Body.String()

	assert.Equal(t, responseRecorder.Code, 400, "Error. Expected code: 400. Actual code: %d", responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Error. Expected: wrong city value. Actual: %s", actualAnswer)

	// Тест №3: Если в параметре `count` указано больше, чем есть всего, должны вернуться все доступные кафе.
	req = httptest.NewRequest("GET", "/cafe?city=moscow&count=7", nil)

	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, req)
	actualCount := strings.Split(responseRecorder.Body.String(), ",")

	assert.Equal(t, totalCount, len(actualCount), "Error. Expected count: %d. Actual code: %d", totalCount, len(actualCount))
}

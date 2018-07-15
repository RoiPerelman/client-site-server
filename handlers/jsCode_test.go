package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
)

func TestEndpoints(t *testing.T) {
	t.Run("test jsCode endpoint", func(t *testing.T) {
		jsonStr := []byte(`{ "jsCode": "console.log('JSCode Test')"}`)
		request, _ := http.NewRequest(http.MethodPost, "/api/user/updateJSCode", bytes.NewBuffer(jsonStr))
		response := httptest.NewRecorder()

		userId := 16
		ctx := context.WithValue(request.Context(), "UserId", userId)
		request = request.WithContext(ctx)

		UpdateJSCode(response, request)

		got := response.Body.String()
		want := "20"
		assertResponseBody(t, got, want)
	})
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

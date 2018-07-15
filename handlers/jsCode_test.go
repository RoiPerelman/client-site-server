package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
	"github.com/roiperelman/client-site-server/models"
	"strings"
	"fmt"
)

type mockDBUserJSStore struct{
	models.DBUserJSStore
}

type mockDb struct {}

func (mockDb *mockDb) UpdateJSCode(id int, jsCode string) error {
	return nil
}

func TestEndpoints(t *testing.T) {
	t.Run("test jsCode endpoint", func(t *testing.T) {
		jsonStr := fmt.Sprintf("{\"jsCode\":\"console.log('JSCode Test')\"}\n")
		request, _ := http.NewRequest(http.MethodPost, "/api/user/updateJSCode", strings.NewReader(jsonStr))
		response := httptest.NewRecorder()

		// add contexts
		userId := 16
		ctx := context.WithValue(request.Context(), "UserId", userId)
		request = request.WithContext(ctx)

		dbUserStoreMock := &mockDBUserJSStore{&mockDb{}}
		ctx = context.WithValue(request.Context(), "DBStore", dbUserStoreMock)
		request = request.WithContext(ctx)

		UpdateJSCode(response, request)

		got := response.Body.String()
		want := "{\"jsCode\":\"console.log('JSCode Test')\"}\n"
		print(want)
		assertResponseBody(t, got, want)
	})
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	println(strings.Compare(got,want))
	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

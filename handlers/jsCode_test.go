package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/roiperelman/client-site-server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// the following type and function are to be used with testify
type mockDBUserJSStoreWithTestify struct {
	mock.Mock
}

func (m *mockDBUserJSStoreWithTestify) UpdateJSCode(id int, jsCode string) error {
	args := m.Called(id, jsCode)
	return args.Error(0)
}

func TestUpdateJSCode(t *testing.T) {
	const UserId = "UserId"
	const DBStore = "DBStore"

	tests := map[string]struct {
		DBUpdateJSCodeResponse error
		contexts               []string
		want                   string
	}{
		"DBStore returns an error": {
			DBUpdateJSCodeResponse: errors.New("mockDBStore error"),
			contexts:               []string{UserId, DBStore},
			want:                   "mockDBStore error\n",
		},
		"UserId doesnt exist in request.context": {
			DBUpdateJSCodeResponse: nil,
			contexts:               []string{DBStore},
			want:                   "update JS code failed\n",
		},
		"DBStore doesnt exist in request.context": {
			DBUpdateJSCodeResponse: nil,
			contexts:               []string{UserId},
			want:                   "db connection failed\n",
		},
		"Successful run": {
			DBUpdateJSCodeResponse: nil,
			contexts:               []string{UserId, DBStore},
			want:                   "{\"jsCode\":\"console.log('JSCode Test')\"}\n",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			jsonStr := fmt.Sprintf("{\"jsCode\":\"console.log('JSCode Test')\"}\n")
			request, _ := http.NewRequest(http.MethodPost, "/api/user/updateJSCode", strings.NewReader(jsonStr))
			response := httptest.NewRecorder()

			for _, testContext := range test.contexts {
				switch testContext {
				case UserId:
					userId := 1
					ctx := context.WithValue(request.Context(), "UserId", userId)
					request = request.WithContext(ctx)
				case DBStore:
					testDBUserJSStore := new(mockDBUserJSStoreWithTestify)
					testDBUserJSStore.
						On("UpdateJSCode", mock.Anything, mock.Anything).
						Return(test.DBUpdateJSCodeResponse)
					ctx := context.WithValue(request.Context(), "DBStore", testDBUserJSStore)
					request = request.WithContext(ctx)
				}
			}

			UpdateJSCode(response, request)

			got := response.Body.String()
			assert.Equal(t, got, test.want, "they should be equal")
		})
	}

	t.Run("Example of a test with mocking the old fashioned way", func(t *testing.T) {
		jsonStr := fmt.Sprintf("{\"jsCode\":\"console.log('JSCode Test')\"}\n")
		request, _ := http.NewRequest(http.MethodPost, "/api/user/updateJSCode", strings.NewReader(jsonStr))
		response := httptest.NewRecorder()

		userId := 1
		ctx := context.WithValue(request.Context(), "UserId", userId)
		request = request.WithContext(ctx)

		dbUserStoreMock := &mockDBUserJSStore{&mockDb{}}
		ctx = context.WithValue(request.Context(), "DBStore", dbUserStoreMock)
		request = request.WithContext(ctx)

		UpdateJSCode(response, request)

		got := response.Body.String()
		want := "{\"jsCode\":\"console.log('JSCode Test')\"}\n"
		assert.Equal(t, got, want, "they should be equal")
	})
}

//the following 2 types and function are used for mocking the old fashioned way
type mockDBUserJSStore struct {
	models.DBUserJSStore
}

type mockDb struct{}

func (mockDb *mockDb) UpdateJSCode(id int, jsCode string) error {
	return nil
}

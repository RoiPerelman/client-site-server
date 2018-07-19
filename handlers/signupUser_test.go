package handlers

import (
	"context"
	"github.com/roiperelman/client-site-server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"
	"fmt"
)

func TestSignupUser(t *testing.T) {
	const DBStore = "DBStore"
	var InputUser = models.User{
		Id: 0,
		Email: "test@test.com",
		Username: "testUsername",
		Password: "testPassword",
		DefaultSection: "",
		PasswordHash: "",
		Token: "",
		IsAuthenticated: false,
		Sections: nil,
		IsMulti: false,
		JSCode: "",
	}
	InputUserJson, _ := json.Marshal(InputUser)

	OutputUser := InputUser
	OutputUser.Id = 1
	OutputUser.Password = ""
	OutputUser.IsAuthenticated = true

	OutputUserJson, _ := json.Marshal(OutputUser)

	tests := map[string]struct {
		InputUserJson []byte
		InsertUserIdResponse  int
		InsertUserErrResponse error
		contexts    []string
		want        string
	}{
		"successful run": {
			InputUserJson: InputUserJson,
			InsertUserIdResponse: 1,
			InsertUserErrResponse: nil,
			contexts: []string{DBStore},
			want:     string(OutputUserJson),
		},
		"DBStore doesnt exist in r.context": {
			InputUserJson: InputUserJson,
			InsertUserIdResponse: 1,
			InsertUserErrResponse: nil,
			contexts: []string{},
			want:     fmt.Sprintln(http.StatusText(http.StatusInternalServerError)),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPost, "/api/user/login", strings.NewReader(string(test.InputUserJson)))
			response := httptest.NewRecorder()

			for _, testContext := range test.contexts {
				switch testContext {
				case DBStore:
					testDBUserContextStore := new(mockDatabaseStore)
					testDBUserContextStore.
						On("InsertUser", mock.Anything, mock.Anything).
						Return(test.InsertUserIdResponse, test.InsertUserErrResponse)
					ctx := context.WithValue(request.Context(), "DBStore", testDBUserContextStore)
					request = request.WithContext(ctx)
				}
			}

			SignupUser(response, request)

			// get response and remove Token (because it is done with a time stamp)
			got := []byte(response.Body.String())
			var gotUser models.User
			err := json.Unmarshal(got, &gotUser)
			if err != nil {
				got = []byte(response.Body.String())
			} else {
				assert.NotEmpty(t, gotUser.Token)
				gotUser.Token = ""
				got, _ = json.Marshal(gotUser)
			}

			assert.Equal(t, string(got), test.want, "they should be equal")
		})
	}
}

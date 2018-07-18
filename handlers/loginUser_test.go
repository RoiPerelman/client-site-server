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

func TestLoginUser(t *testing.T) {
	const DBStore = "DBStore"
	var TestUser = models.User{
		Id: 1,
		Email: "test@test.com",
		Username: "testUsername",
		Password: "",
		DefaultSection: "1",
		PasswordHash: "$2a$14$WwH3itw0bjE5FuR3F40eHOH0rg1FY6CoXZEbAPot55b0umBHz9hD6",
		Token: "",
		IsAuthenticated: true,
		Sections: map[string]models.Section{"1": {1, "1", "mockSection",
			models.Contexts{[]string{}, []string{}, []string{}}}},
		IsMulti: false,
		JSCode: "",
	}
	TestUserJson, _ := json.Marshal(TestUser)

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

	InputUserJsonWrongPassword := InputUser
	InputUserJsonWrongPassword.Password = "wrongPassword"
	InputUserJsonWrongPasswordJson, _ := json.Marshal(InputUserJsonWrongPassword)

	tests := map[string]struct {
		InputUserJson []byte
		GetUserByEmail *models.User
		contexts                   []string
		want                       string
	}{
		"successful run": {
			InputUserJson: InputUserJson,
			GetUserByEmail: &TestUser,
			contexts: []string{DBStore},
			want:     string(TestUserJson),
		},
		"DBStore doesnt exist in r.context": {
			InputUserJson: InputUserJson,
			GetUserByEmail: &TestUser,
			contexts: []string{},
			want:     fmt.Sprintln(http.StatusText(http.StatusInternalServerError)),
		},
		"DBStore returns nil dbUser": {
			InputUserJson: InputUserJsonWrongPasswordJson,
			GetUserByEmail: nil,
			contexts: []string{DBStore},
			want:     fmt.Sprintln(http.StatusText(http.StatusUnauthorized)),
		},
		"InputUserWrongPassword": {
			InputUserJson: InputUserJsonWrongPasswordJson,
			GetUserByEmail: &TestUser,
			contexts: []string{DBStore},
			want:     fmt.Sprintln(http.StatusText(http.StatusUnauthorized)),
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
						On("GetUserByEmail", mock.Anything).Return(&TestUser)
					ctx := context.WithValue(request.Context(), "DBStore", testDBUserContextStore)
					request = request.WithContext(ctx)
				}
			}

			LoginUser(response, request)

			// get response and remove Token (because it is done with a time stamp)
			got := []byte(response.Body.String())
			var gotUser models.User
			err := json.Unmarshal(got, &gotUser)
			if err != nil {
				got = []byte(response.Body.String())
			} else {
				gotUser.Token = ""
				got, _ = json.Marshal(gotUser)
			}

			assert.Equal(t, string(got), test.want, "they should be equal")
		})
	}
}

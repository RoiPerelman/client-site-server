package handlers

import (
	"context"
	"fmt"
	"github.com/roiperelman/client-site-server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestContexts(t *testing.T) {
	const DBStore = "DBStore"

	tests := map[string]struct {
		GetUserSectionBySectionsId models.Section
		contexts                   []string
		want                       string
	}{
		"AddContextItem successful run": {
			GetUserSectionBySectionsId: models.Section{1, "1", "mockSection",
				models.Contexts{[]string{}, []string{}, []string{}}},
			contexts: []string{DBStore},
			want:     "{\"id\":1,\"sectionId\":\"1\",\"name\":\"mockSection\",\"contexts\":{\"product\":[],\"cart\":[],\"category\":[]}}\n",
		},
		"AddContextItem without r.context": {
			GetUserSectionBySectionsId: models.Section{1, "1", "mockSection",
				models.Contexts{[]string{}, []string{}, []string{}}},
			contexts: []string{},
			want:     fmt.Sprintln(http.StatusText(http.StatusInternalServerError)),
		},
		"DelContextItem": {
			GetUserSectionBySectionsId: models.Section{1, "1", "mockSection",
				models.Contexts{[]string{}, []string{}, []string{}}},
			contexts: []string{DBStore},
			want:     "{\"id\":1,\"sectionId\":\"1\",\"name\":\"mockSection\",\"contexts\":{\"product\":[],\"cart\":[],\"category\":[]}}\n",
		},
		"DelContextItem without r.context": {
			GetUserSectionBySectionsId: models.Section{1, "1", "mockSection",
				models.Contexts{[]string{}, []string{}, []string{}}},
			contexts: []string{},
			want:     "{\"id\":1,\"sectionId\":\"1\",\"name\":\"mockSection\",\"contexts\":{\"product\":[],\"cart\":[],\"category\":[]}}\n",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			jsonStr := fmt.Sprintf(`{"sectionsId":1,"sectionId":"1","contextType":"PRODUCT","item":"1"}`)
			request, _ := http.NewRequest(http.MethodPost, "/api/user/addContextItem", strings.NewReader(jsonStr))
			response := httptest.NewRecorder()

			for _, testContext := range test.contexts {
				switch testContext {
				case DBStore:
					testDBUserContextStore := new(mockDatabaseStore)
					testDBUserContextStore.
						On("AddContextTypeItem", mock.Anything).Return()
					testDBUserContextStore.
						On("DelContextTypeItem", mock.Anything).Return()
					testDBUserContextStore.
						On("GetUserSectionBySectionsId", mock.Anything).
						Return(test.GetUserSectionBySectionsId)
					ctx := context.WithValue(request.Context(), "DBStore", testDBUserContextStore)
					request = request.WithContext(ctx)
				}
			}

			AddContextItem(response, request)

			got := response.Body.String()
			assert.Equal(t, got, test.want, "they should be equal")
		})
	}
}

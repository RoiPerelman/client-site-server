package handlers

import (
	"github.com/stretchr/testify/mock"
	"github.com/roiperelman/client-site-server/models"
)

// the following type and function are to be used with testify
type mockDatabaseStore struct {
	mock.Mock
}

// DBUserJSStore
func (m *mockDatabaseStore) UpdateJSCode(id int, jsCode string) error {
	args := m.Called(id, jsCode)
	return args.Error(0)
}

// DBUserStore
func (m *mockDatabaseStore) InsertUser(user *models.User) (int, error) {
	args := m.Called(user)
	return args.Get(0).(int), args.Error(1)
}
func (m *mockDatabaseStore) GetUserById(id int) *models.User {
	args := m.Called(id)
	return args.Get(0).(*models.User)
}
func (m *mockDatabaseStore) GetUserByEmail(email string) *models.User {
	args := m.Called(email)
	return args.Get(0).(*models.User)
}
func (m *mockDatabaseStore) GetUserByUsername(username string) *models.User {
	args := m.Called(username)
	return args.Get(0).(*models.User)
}

// DBSectionStore
func (m *mockDatabaseStore) GetAllUserIdSections(userId int) map[string]models.Section {
	args := m.Called(userId)
	return args.Get(0).(map[string]models.Section)
}
func (m *mockDatabaseStore) GetUserSectionBySectionsId(sectionsId int) models.Section {
	args := m.Called(sectionsId)
	return args.Get(0).(models.Section)
}
func (m *mockDatabaseStore) AddSection(userId int, section models.Section) int {
	args := m.Called(userId, section)
	return args.Get(0).(int)
}
func (m *mockDatabaseStore) DelSection(id int, section models.Section) {
	 m.Called(id, section)
	return
}
func (m *mockDatabaseStore) UpdateIsMultipleSectionFeature(id int, isMulti bool) {
	m.Called(id)
	return
}

// DBContextsStore
func (m *mockDatabaseStore) GetContextsBySectionsId(sectionsIdentifier int) models.Contexts {
	args := m.Called(sectionsIdentifier)
	return args.Get(0).(models.Contexts)
}
func (m *mockDatabaseStore) AddContextTypeItem(contextItem *models.ContextItem) {
	m.Called(contextItem)
	return
}
func (m *mockDatabaseStore) DelContextTypeItem(contextItem *models.ContextItem) {
	m.Called(contextItem)
	return
}

package validateuser_test

import (
	"errors"
	"testbadry/validateuser"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockWithTestifyAPI struct {
	mock.Mock
}

func (m *mockWithTestifyAPI) GetUserInfo(user string) (*validateuser.UserInformation, error) {
	args := m.Called(user)
	userInfo := (*validateuser.UserInformation)(nil)
	if args.Get(0) != nil {
		userInfo = args.Get(0).(*validateuser.UserInformation)
	} else {
		userInfo = nil
	}
	return userInfo, args.Error(1)
}

func Test_ValidateUser_fail_with_user_not_found_TS(t *testing.T) {
	api := new(mockWithTestifyAPI)
	api.On("GetUserInfo", "kaweelhandsome").Return(nil, errors.New("Not Found"))
	s := validateuser.NewValidateUserService(api)
	resp, err := s.ValidateUser("kaweelhandsome", mock.Anything)
	if assert.Error(t, err) {
		assert.False(t, resp)
		assert.Equal(t, "Not Found", err.Error())
	}
	api.AssertExpectations(t)
	api.AssertNumberOfCalls(t, "GetUserInfo", 1)
}

func Test_ValidateUser_success_with_name_not_match_TS(t *testing.T) {
	userInfo := &validateuser.UserInformation{
		Name: "Mongkol Tontan",
	}
	api := new(mockWithTestifyAPI)
	api.On("GetUserInfo", "kaweel").Return(userInfo, nil)
	s := validateuser.NewValidateUserService(api)
	resp, err := s.ValidateUser("kaweel", "Kawee Lertrungmongkol")
	if assert.NoError(t, err) {
		assert.False(t, resp)
	}
	api.AssertExpectations(t)
	api.AssertNumberOfCalls(t, "GetUserInfo", 1)
}

func Test_ValidateUser_success_with_name_match_TS(t *testing.T) {
	userInfo := &validateuser.UserInformation{
		Name: "Kawee Lertrungmongkol",
	}
	api := new(mockWithTestifyAPI)
	api.On("GetUserInfo", "kaweel").Return(userInfo, nil)
	s := validateuser.NewValidateUserService(api)
	resp, err := s.ValidateUser("kaweel", "Kawee Lertrungmongkol")
	if assert.NoError(t, err) {
		assert.True(t, resp)
	}
	api.AssertExpectations(t)
	api.AssertNumberOfCalls(t, "GetUserInfo", 1)
}

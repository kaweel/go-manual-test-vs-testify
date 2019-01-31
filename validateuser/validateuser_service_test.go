package validateuser_test

import (
	"errors"
	"reflect"
	"testbadry/validateuser"
	"testing"
)

type mockAPI struct {
	UserInformation *validateuser.UserInformation
	Err             error
}

func (m *mockAPI) SetResponse(userInformation *validateuser.UserInformation, err error) {
	m.UserInformation = userInformation
	m.Err = err
}

func (m *mockAPI) GetUserInfo(user string) (*validateuser.UserInformation, error) {
	return m.UserInformation, m.Err
}

func Test_ValidateUser_fail_with_user_not_found(t *testing.T) {
	api := &mockAPI{}
	api.SetResponse(nil, errors.New("Not Found"))
	s := validateuser.NewValidateUserService(api)
	resp, err := s.ValidateUser("kaweelhandsome", "")
	if !reflect.DeepEqual(err, nil) {
		if !reflect.DeepEqual("Not Found", err.Error()) {
			t.Errorf("Expected error msg %v but got %v\n", nil, err)
		}
		if !reflect.DeepEqual(resp, false) {
			t.Errorf("Expected %v but got %v\n", false, resp)
		}
	}
}

func Test_ValidateUser_success_with_name_not_match(t *testing.T) {
	userInfo := &validateuser.UserInformation{
		Name: "Mongkol Tontan",
	}
	api := &mockAPI{}
	api.SetResponse(userInfo, nil)
	s := validateuser.NewValidateUserService(api)
	resp, err := s.ValidateUser("kaweel", "Kawee Lertrungmongkol")
	if reflect.DeepEqual(err, nil) {
		if !reflect.DeepEqual(resp, false) {
			t.Fatalf("Expected %v but got %v", false, resp)
		}
	}
}

func Test_ValidateUser_success_with_name_match(t *testing.T) {
	userInfo := &validateuser.UserInformation{
		Name: "Kawee Lertrungmongkol",
	}
	api := &mockAPI{}
	api.SetResponse(userInfo, nil)
	s := validateuser.NewValidateUserService(api)
	resp, err := s.ValidateUser("kaweel", "Kawee Lertrungmongkol")
	if reflect.DeepEqual(err, nil) {
		if !reflect.DeepEqual(resp, true) {
			t.Fatalf("Expected %v but got %v", true, resp)
		}
	}
}

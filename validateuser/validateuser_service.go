package validateuser

type ValidateUserService interface {
	ValidateUser(user string, name string) (bool, error)
}

type service struct {
	api GithubAPI
}

func NewValidateUserService(api GithubAPI) ValidateUserService {
	return &service{api}
}

func (s *service) ValidateUser(user string, name string) (bool, error) {
	res, err := s.api.GetUserInfo(user)
	if err != nil {
		return false, err
	}
	if res.Name != name {
		return false, err
	}
	return true, nil
}

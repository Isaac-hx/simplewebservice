// File to implement service user,handling to request and response from user
package repository

type UserService struct {
	repo UserRepository
}

func (s *UserService) ExecuteGetUserById(id int) (string, error) {
	return s.repo.GetUserById(id)
}

// file to implement mock object
package repository

// mock_repository.go
type MockUserRepository struct {
	//jika properti ini tidak diiniasilasasi,maka yang terjadi adalah default value dari tipe ini adalah nil
	GetUserByIDFunc func(int) (string, error) //callback
	Calls           map[string]int
}

// Method getuserbyid implement interface userRepository
func (m *MockUserRepository) GetUserById(id int) (string, error) {
	m.Calls["GetUserByID"]++
	//pemanggilan callback pada object MockUserRepository
	return m.GetUserByIDFunc(id)
}

// Constructor untuk mockUserRepository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		//Menginisialiasi properti calls
		Calls: make(map[string]int),
	}
}

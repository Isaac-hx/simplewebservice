// File to testing mock RealUserRepository
package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_Success(t *testing.T) {
	//Menginisialisasi object MockUserRepository()
	mockRepo := NewMockUserRepository()
	//Insert value callback from object mock
	mockRepo.GetUserByIDFunc = func(id int) (string, error) {
		return "Mock User", nil

	}

	//Menginisialaisi object userService
	//repo property adalah tipe interface userRepository
	//object userservice
	service := &UserService{repo: mockRepo}
	result, err := service.ExecuteGetUserById(1)
	assert.NoError(t, err)
	assert.Equal(t, "Mock User", result)
	assert.Equal(t, 1, mockRepo.Calls["GetUserById"])
}

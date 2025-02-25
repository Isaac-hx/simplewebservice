package repository

// Membuat depedency untuk dimock
type UserRepository interface {
	GetUserById(int) (string, error)
}

// file to implement object to database!
package repository

type RealUserRepository struct{}

func (r *RealUserRepository) GetUserById(id int) (string, error) {
	return "real user", nil

}

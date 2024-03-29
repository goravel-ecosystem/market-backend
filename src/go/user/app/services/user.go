package services

import (
	"market.goravel.dev/user/app/models"
)

type User interface {
	GetUserByEmail(email string) (*models.User, error)
	IsEmailExist(email string) (bool, error)
	Register(name, email, password string) (*models.User, error)
}

type UserImpl struct {
	userModel models.UserInterface
}

func NewUserImpl() *UserImpl {
	return &UserImpl{
		userModel: models.NewUser(),
	}
}

func (r *UserImpl) GetUserByEmail(email string) (*models.User, error) {
	return r.userModel.GetUserByEmail(email, []string{"id", "name", "email", "password", "avatar", "summary"})
}

func (r *UserImpl) IsEmailExist(email string) (bool, error) {
	user, err := r.userModel.GetUserByEmail(email, []string{"id"})
	if err != nil {
		return false, err
	}

	return user.ID > 0, nil
}

func (r *UserImpl) Register(name, email, password string) (*models.User, error) {
	return r.userModel.Register(name, email, password)
}

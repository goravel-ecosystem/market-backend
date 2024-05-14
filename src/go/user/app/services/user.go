package services

import (
	"context"

	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"

	protouser "market.goravel.dev/proto/user"
	"market.goravel.dev/user/app/models"
	utilerrors "market.goravel.dev/utils/errors"
)

type User interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUsers(ids []string) ([]*models.User, error)
	IsEmailExist(email string) (bool, error)
	Register(name, email, password string) (*models.User, error)
	UpdateUser(ctx context.Context, req *protouser.UpdateUserRequest) (*models.User, error)
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

func (r *UserImpl) GetUserByID(id string) (*models.User, error) {
	return r.userModel.GetUserByID(id, []string{"id", "name", "avatar", "summary"})
}

func (r *UserImpl) GetUsers(ids []string) ([]*models.User, error) {
	return r.userModel.GetUsers(ids, []string{"id", "name", "avatar"})
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

func (r *UserImpl) UpdateUser(ctx context.Context, req *protouser.UpdateUserRequest) (*models.User, error) {
	user, err := r.userModel.GetUserByID(req.GetId(), []string{})
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, utilerrors.NewNotFound(facades.Lang(ctx).Get("not_exist.user"))
	}

	if user.ID != cast.ToUint64(req.GetUserId()) {
		return nil, utilerrors.NewUnauthorized(facades.Lang(ctx).Get("forbidden.update_user"))
	}

	user.Name = req.GetName()
	user.Avatar = req.GetAvatar()
	user.Summary = req.GetSummary()

	password := req.GetPassword()
	if password != "" {
		hashedPassword, err := facades.Hash().Make(req.GetPassword())
		if err != nil {
			return nil, utilerrors.NewInternalServerError(err)
		}
		user.Password = hashedPassword
	}

	if err := r.userModel.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"

	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
)

type UserInterface interface {
	GetUserByEmail(email string, fields []string) (*User, error)
	GetUserByID(id string, fields []string) (*User, error)
	GetUsers(ids []string, fields []string) ([]*User, error)
	Register(name, email, password string) (*User, error)
	UpdateUser(user *User) error
}

type User struct {
	UUIDModel
	Email    string
	Password string
	Name     string
	Avatar   string
	Summary  string
	orm.SoftDeletes
}

func NewUser() *User {
	return &User{}
}

func (r *User) GetUserByEmail(email string, fields []string) (*User, error) {
	var user User
	if err := facades.Orm().Query().Where("email", email).Select(fields).First(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *User) GetUserByID(id string, fields []string) (*User, error) {
	var user User

	query := facades.Orm().Query().Where("id", id)

	if len(fields) > 0 {
		query = query.Select(fields)
	}

	if err := query.First(&user); err != nil {
		return nil, utilserrors.NewInternalServerError(err)
	}

	return &user, nil
}

func (r *User) GetUsers(ids []string, fields []string) ([]*User, error) {
	var users []*User

	var userIDs []any
	for _, id := range ids {
		userIDs = append(userIDs, id)
	}

	if err := facades.Orm().Query().WhereIn("id", userIDs).Select(fields).Find(&users); err != nil {
		return nil, utilserrors.NewInternalServerError(err)
	}

	return users, nil
}

func (r *User) Register(name, email, password string) (*User, error) {
	hashedPassword, err := facades.Hash().Make(password)
	if err != nil {
		return nil, err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	user.ID = r.GetID()
	if err := facades.Orm().Query().Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *User) ToProto() *protouser.User {
	return &protouser.User{
		Id:      cast.ToString(r.ID),
		Name:    r.Name,
		Email:   r.Email,
		Avatar:  r.Avatar,
		Summary: r.Summary,
	}
}

func (r *User) UpdateUser(user *User) error {
	if err := facades.Orm().Query().Save(user); err != nil {
		return utilserrors.NewInternalServerError(err)
	}

	return nil
}

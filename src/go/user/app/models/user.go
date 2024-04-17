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
	Register(name, email, password string) (*User, error)
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
	if err := facades.Orm().Query().Where("id", id).Select(fields).First(&user); err != nil {
		return nil, utilserrors.NewInternalServerError(err)
	}

	return &user, nil
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

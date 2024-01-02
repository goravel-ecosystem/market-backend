package services

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/contracts/translation"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"

	"github.com/goravel-ecosystem/market-backend/user/app/helper"
)

//go:generate mockery --name=Notification
type Notification interface {
	SendEmailRegisterCode(ctx context.Context, email string) (key string, err error)
	VerifyEmailRegisterCode(key, code string) bool
}

type NotificationImpl struct {
}

func NewNotificationImpl() *NotificationImpl {
	return &NotificationImpl{}
}

func (r *NotificationImpl) SendEmailRegisterCode(ctx context.Context, email string) (key string, err error) {
	var code int
	if helper.IsProduction() || helper.IsDevelopment() {
		code = rand.Intn(899999) + 100000
	} else {
		code = 123123
	}

	key = r.getEmailRegisterCodeKey(email)
	if err := facades.Cache().Put(key, fmt.Sprintf("%d", code), 60*5*time.Second); err != nil {
		return "", err
	}

	if helper.IsProduction() || helper.IsDevelopment() {
		subject, _ := facades.Lang(ctx).Get("register_code.subject", translation.Option{
			Replace: map[string]string{
				"code": fmt.Sprintf("%d", code),
			},
		})
		html, _ := facades.Lang(ctx).Get("register_code.content", translation.Option{
			Replace: map[string]string{
				"code": fmt.Sprintf("%d", code),
			},
		})
		if err := facades.Mail().To([]string{email}).Content(mail.Content{
			Subject: subject,
			Html:    html,
		}).Queue(); err != nil {
			return "", err
		}
	}

	return key, nil
}

func (r *NotificationImpl) VerifyEmailRegisterCode(key, code string) bool {
	if facades.Cache().GetString(key) == code {
		facades.Cache().Forget(key)

		return true
	}

	return false
}

func (r *NotificationImpl) getEmailRegisterCodeKey(email string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("email_register_code_%s_%s", email, carbon.Now().ToDateNanoString()))))
}

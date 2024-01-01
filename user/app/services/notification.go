package services

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"

	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/contracts/translation"
	"github.com/goravel/framework/facades"
)

type Notification interface {
	SendRegisterEmailCode(ctx context.Context, email string) error
	VerifyRegisterEmailCode(email, code string) bool
}

type NotificationImpl struct {
}

func NewNotification() *NotificationImpl {
	return &NotificationImpl{}
}

func (r *NotificationImpl) SendRegisterEmailCode(ctx context.Context, email string) error {
	code := rand.Intn(899999) + 100000
	if err := facades.Cache().Put(r.getRegisterEmailCodeKey(email), fmt.Sprintf("%d", code), 60*5); err != nil {
		return err
	}

	if IsProduction() {
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
			return err
		}
	}

	return nil
}

func (r *NotificationImpl) VerifyRegisterEmailCode(email, code string) bool {
	if facades.Cache().GetString(r.getRegisterEmailCodeKey(email)) == code {
		facades.Cache().Forget(r.getRegisterEmailCodeKey(email))

		return true
	}

	return false
}

func (r *NotificationImpl) getRegisterEmailCodeKey(email string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(email)))
}

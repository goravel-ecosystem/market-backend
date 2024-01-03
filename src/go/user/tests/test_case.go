package tests

import (
	"github.com/goravel/framework/testing"

	"market.goravel.dev/user/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}

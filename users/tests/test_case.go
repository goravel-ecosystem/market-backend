package tests

import (
	"github.com/goravel/framework/testing"

	"github.com/goravel-ecosystem/market-backend/users/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}

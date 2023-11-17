package tests

import (
	"github.com/goravel/framework/testing"

	"github.com/goravel-ecosystem/market-backend/gateway/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}

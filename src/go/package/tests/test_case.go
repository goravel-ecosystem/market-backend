package tests

import (
	"github.com/goravel/framework/testing"

	"market.goravel.dev/package/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}

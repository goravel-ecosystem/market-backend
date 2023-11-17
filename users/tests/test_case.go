package tests

import (
	"github.com/goravel/framework/testing"

	"users/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}

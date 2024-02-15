package interceptors

import (
	"context"
	"errors"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/testing/mock"
	"google.golang.org/grpc"

	"market.goravel.dev/proto/base"
	utilserrors "market.goravel.dev/utils/errors"
)

func TestResponse(t *testing.T) {
	beforeEach := func() {
		mockFactory := mock.Factory()
		mockFactory.Log()
	}

	tests := []struct {
		name           string
		req            any
		handler        grpc.UnaryHandler
		expectResponse any
		expectError    error
	}{
		{
			name: "Happy path",
			req:  "Goravel",
			handler: func(ctx context.Context, req any) (any, error) {
				return "Hello, " + req.(string), nil
			},
			expectResponse: "Hello, Goravel",
		},
		{
			name: "Sad path, handler returns utilserrors.ErrorWithCode",
			req:  "Goravel",
			handler: func(ctx context.Context, req any) (any, error) {
				return nil, utilserrors.NewValidate("error")
			},
			expectResponse: &base.Response{
				Status: &base.Status{
					Code:  http.StatusBadRequest,
					Error: "error",
				},
			},
		},
		{
			name: "Sad path, handler returns unknown error",
			req:  "Goravel",
			handler: func(ctx context.Context, req any) (any, error) {
				return nil, errors.New("error")
			},
			expectResponse: &base.Response{
				Status: &base.Status{
					Code:  http.StatusInternalServerError,
					Error: "Internal server error, please try it later.",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			resp, err := Response()(context.Background(), test.req, nil, test.handler)
			assert.Equal(t, test.expectResponse, resp)
			assert.Equal(t, test.expectError, err)
		})
	}
}

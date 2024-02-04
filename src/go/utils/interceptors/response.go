package interceptors

import (
	"context"
	"errors"

	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"google.golang.org/grpc"

	"market.goravel.dev/proto/base"
	utilserrors "market.goravel.dev/utils/errors"
)

func Response() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			var errorWithCode utilserrors.ErrorWithCode
			if errors.As(err, &errorWithCode) {
				return &base.Response{
					Status: &base.Status{
						Code:  int32(errorWithCode.Code()),
						Error: errorWithCode.Error(),
					},
				}, nil
			}

			facades.Log().WithContext(ctx).With(map[string]any{
				"req": req,
			}).Error(err)

			return &base.Response{
				Status: &base.Status{
					Code:  contractshttp.StatusInternalServerError,
					Error: "Internal server error, please try it later.",
				},
			}, nil
		}

		return resp, err
	}
}

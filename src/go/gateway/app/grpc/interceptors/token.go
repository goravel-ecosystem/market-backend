package interceptors

import (
	"context"
	"net/http"

	"google.golang.org/protobuf/proto"
)

type TokenResponse interface {
	GetToken() string
}

func Token(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	if tokenResponse, ok := resp.(TokenResponse); ok && tokenResponse.GetToken() != "" {
		w.Header().Set("Authorization", tokenResponse.GetToken())
	}

	return nil
}

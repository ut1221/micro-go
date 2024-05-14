package middleware

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"micro/internal/pkg"
	"strings"

	middleware2 "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func Auth() middleware2.Middleware {
	return func(handler middleware2.Handler) middleware2.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				auth := tr.RequestHeader().Get("Authorization")
				if auth == "" {
					return nil, errors.New("no Auth")
				}
				splitStr := strings.Split(auth, " ")
				token := splitStr[len(splitStr)-1]
				userClaims, err := pkg.AnalyseToken(token)
				if err != nil {
					return nil, err
				}
				if userClaims.Identity == "" {
					return nil, errors.New("no Auth")
				}
				ctx = metadata.NewServerContext(ctx, metadata.New(map[string][]string{
					"username": []string{userClaims.Name},
					"identity": []string{userClaims.Identity},
				}))
			}
			return handler(ctx, req)
		}
	}
}

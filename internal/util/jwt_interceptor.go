package util

import (
	"booking-identity-management/internal/adapter"
	"booking-identity-management/internal/common"
	"context"
	"regexp"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	patternStr = "^Bearer (\\S*)"
)

func checkContains(arr []string, in string) bool {
	for _, item := range arr {
		if strings.Contains(in, item) {
			return true
		}
	}
	return false
}

func JWTAuthenticationInterceptor(jwtAdapter adapter.JWTAdapter, excludePaths []string) grpc.UnaryServerInterceptor {
	pattern, _ := regexp.Compile(patternStr)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}

		// excluded path need not handle jwt
		if checkContains(excludePaths, info.FullMethod) {
			return handler(ctx, req)
		}

		if authorization, ok := md["authorization"]; ok && len(authorization) != 0 {
			firstItem := authorization[0]
			if firstItem != "" && pattern.Match([]byte(firstItem)) {
				tokens := pattern.FindSubmatch([]byte(firstItem))
				if len(tokens) == 2 && string(tokens[1]) != "" {
					jwt := string(tokens[1])
					claim, err := jwtAdapter.VerifyToken(ctx, jwt, "bim", "")
					if err == nil {
						md[common.DomainMDKey] = []string{claim.Domain}
						ctx = metadata.NewIncomingContext(ctx, md)
						return handler(ctx, req)
					}
					// reason := common.ParseError(err)
					// switch reason {
					// case common.ReasonJWTExpired:
					// if claim.RefreshToken != "" {
					// 	newJWT, err := userActionUsecase.RefreshToken(ctx,
					// 		dto.UserBaseDTO{RefreshToken: claim.RefreshToken, UserDomain: claim.Domain, JWTToken: jwt})
					// 	if err != nil {
					// 		// Log here
					// 		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
					// 	}
					// 	if newJWT != "" {
					// 		newMD := map[string][]string{
					// 			"authorization": {newJWT},
					// 		}
					// 		_ = grpc.SetHeader(ctx, newMD)
					// 		md[common.DomainMDKey] = []string{claim.Domain}
					// 		return handler(ctx, req)
					// 	}
					return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
					// }
					// default:
					// 	return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
					// }
				}
			}
		}
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}
}

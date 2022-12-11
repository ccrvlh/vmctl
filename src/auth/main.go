package auth

// import (
// 	"context"
// 	"encoding/base64"
// )

// type basicAuth struct {
// 	token string
// }

// func basic(t string) basicAuth {
// 	return basicAuth{token: t}
// }

// func (b basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
// 	enc := base64.StdEncoding.EncodeToString([]byte(b.token))

// 	return map[string]string{
// 		"authorization": "Basic " + enc,
// 	}, nil
// }

// func (basicAuth) RequireTransportSecurity() bool {
// 	return false
// }

// package dialler

// import (
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// // New process the dial config and returns a grpc.ClientConn. The caller is
// // responsible for closing the connection.
// func New(address, basicAuthToken string) (*grpc.ClientConn, error) {
// 	dialOpts := []grpc.DialOption{
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	}

// 	if basicAuthToken != "" {
// 		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(
// 			basic(basicAuthToken),
// 		))
// 	}

// 	return grpc.Dial(
// 		address,
// 		dialOpts...,
// 	)
// }

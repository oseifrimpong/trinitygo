package grpc

import (
	"context"
	"sync"

	helloworldpb "github.com/PolarPanda611/trinitygo/example/pb/helloworld"

	"github.com/PolarPanda611/trinitygo/example/server/domain/service"

	"github.com/PolarPanda611/trinitygo"
)

func init() {
	trinitygo.BindController("helloworld.Greeter", &sync.Pool{
		New: func() interface{} {
			service := new(Server)
			return service
		},
	})
}

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	UserService service.UserService
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	// md, ok := metadata.FromIncomingContext(ctx)
	name := s.UserService.GetUserNameByID(in.GetName())
	return &helloworldpb.HelloReply{Message: name}, nil
}

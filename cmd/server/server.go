package main

import (
	"context"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/go-leo/leo"
	"github.com/go-leo/leo/log"
	middlewarecontext "github.com/go-leo/leo/middleware/context"
	middlewarelog "github.com/go-leo/leo/middleware/log"
	"github.com/go-leo/leo/middleware/recovery"
	"github.com/go-leo/leo/middleware/requestid"
	v1 "github.com/xiaoyan648/project-layout/api/helloworld/v1"
	"github.com/xiaoyan648/project-layout/internal/pkg/conf"
	"github.com/xiaoyan648/project-layout/internal/service"
	"google.golang.org/grpc"
)

func NewHttpServer(conf *conf.Server, server *service.GreeterService, logger log.Logger) *leo.HttpOptions {
	return &leo.HttpOptions{
		Port: conf.Http.Port,
		GinMiddlewares: []gin.HandlerFunc{
			requestid.GinMiddleware(),
			middlewarecontext.GinMiddleware(func(ctx context.Context) context.Context {
				traceID, _ := requestid.FromContext(ctx)

				return log.NewContext(ctx, logger.Clone().With(log.F{K: "TraceID", V: traceID}))
			}),
			middlewarelog.GinMiddleware(log.FromContextOrDiscard),
			recovery.GinMiddleware(func(ctx *gin.Context, a any) {
				logger := log.FromContextOrDiscard(ctx.Request.Context())
				logger.Errorf("%s %+v", "[Panic]", a)
				debug.PrintStack()
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}),
		},
		Routers: v1.GreeterAPIRouters(server),
	}
}

// TODO grpc 兼容
func NewGrpcServer(conf *conf.Server, server *service.GreeterService, logger log.Logger) (*grpc.Server, *leo.GRPCOptions) {
	var grpcOpts = []grpc.ServerOption{}
	srv := grpc.NewServer(grpcOpts...)
	v1.RegisterGreeterServer(srv, server)
	// lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", conf.Grpc.Port))
	// if err != nil {
	// 	panic(err)
	// }
	// panic(srv.Serve(lis))
	opts := &leo.GRPCOptions{
		Port:              conf.Grpc.Port,
		GRPCServerOptions: grpcOpts,
	}
	return srv, opts
}

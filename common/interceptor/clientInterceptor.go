package interceptor

import (
	"context"
	// "time"
	// "log"
	"google.golang.org/grpc"
	"TDS-backend/common/errorx"
	"google.golang.org/grpc/status"
)

func ClientErrorinterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	// log.Printf("method == %s ; req == %v ; rep == %v ; duration == %s ; error == %v\n", method, req, reply, time.Since(start), err)
	if err != nil {
		err = errorx.NewCodeError(int(status.Code(err)), err.Error())
	}
	return err
 }
 
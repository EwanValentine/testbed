package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	greeter "github.com/EwanValentine/testbed/services/greeter/gen/go/proto"
)

func main() {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}))

	ctx := context.Background()
	var opts []grpc.DialOption

	// I know it's deprecated, but can't be arsed figuring out what it is now
	opts = append(opts, grpc.WithInsecure())
	if err := greeter.RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts); err != nil {
		log.Panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}

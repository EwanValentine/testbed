package gateway

import (
	"context"
	greeter "github.com/EwanValentine/testbed/services/greeter/gen/go/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
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
	if err := greeter.RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts); err != nil {
		log.Panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}

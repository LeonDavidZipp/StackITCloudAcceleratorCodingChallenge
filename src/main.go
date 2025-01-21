package main

import (
	"context"
	"fmt"
	"github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api"
	s "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/gen/forwarder_service"
	sgrpcpb "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/gen/grpc/forwarder_service/pb"
	sgrpc "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/gen/grpc/forwarder_service/server"
	shttp "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/gen/http/forwarder_service/server"
	c "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/config"
	goalog "goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// Creates a new HTTP server
func newHttpServer(ctx context.Context, ep *s.Endpoints) *http.Server {
	// prepare server parameters
	mux := goahttp.NewMuxer()
	decoder := goahttp.RequestDecoder
	encoder := goahttp.ResponseEncoder
	errorHandler := func(ctx context.Context, w http.ResponseWriter, err error) {
		log.Printf("encoding error: %s", err.Error())
	}

	server := shttp.New(ep, mux, decoder, encoder, errorHandler, nil)
	server.Mount(mux)

	var handler http.Handler = mux
	handler = goalog.HTTP(ctx)(handler)

	httpServer := &http.Server{
		Addr:              ":" + c.HTTP_SERVER_PORT,
		Handler:           handler,
		ReadHeaderTimeout: 60 * time.Second,
	}

	return httpServer
}

// Creates a new gRPC server
func newGrpcServer(ep *s.Endpoints) (*grpc.Server, net.Listener, error) {
	server := sgrpc.New(ep, nil)

	grpcServer := grpc.NewServer()
	sgrpcpb.RegisterForwarderServiceServer(grpcServer, server)

	listener, err := net.Listen(
		"tcp",
		":"+c.GRPC_SERVER_PORT,
	)
	if err != nil {
		return nil, nil, err
	}

	return grpcServer, listener, nil
}

func main() {
	// prepare context
	ctx := goalog.Context(
		context.Background(),
		goalog.WithDebug(),
	)
	ctx, cancel := context.WithTimeout(
		ctx,
		60*time.Second,
	)
	defer cancel()

	// prepare service
	service := api.NewForwarderService(c.TELEGRAM_CHAT_ID)
	endpoints := s.NewEndpoints(service)

	// create servers
	httpServer := newHttpServer(ctx, endpoints)

	grpcServer, grpcListener, err := newGrpcServer(endpoints)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	// start http server
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("http server running on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// start grpc server
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("grpc server running on %s\n", grpcListener.Addr())
		if err := grpcServer.Serve(grpcListener); err != nil {
			panic(err)
		}
	}()

	wg.Wait()

	if err := httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}

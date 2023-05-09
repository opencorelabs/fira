package application

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

func (a *App) StartGRPC(ctx context.Context) error {
	log := a.Logger().Sugar().Named("startup")
	svc := &api.FiraApiService{}
	grpcServer := grpc.NewServer()
	v1.RegisterFiraServiceServer(grpcServer, svc)

	// start the GRPC service
	a.StartService(ctx, "grpc-server", func(ctx context.Context, errChan chan error) Finalizer {
		go func() {
			defer func() {
				if p := recover(); p != nil {
					errChan <- fmt.Errorf("recovered from panic: %#v", p)
				}
			}()

			listener, listenErr := net.Listen("tcp", a.cfg.GrpcUrl)
			if listenErr != nil {
				errChan <- fmt.Errorf("failed to listen: %w", listenErr)
			}

			log.Infow("gRPC server ready", "addr", a.cfg.GrpcUrl)

			errChan <- grpcServer.Serve(listener)
		}()

		return func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		}
	})

	// dial the gRPC server above to make a client connection
	conn, connErr := grpc.Dial(a.cfg.GrpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if connErr != nil {
		return fmt.Errorf("fail to dial grpc server: %v", connErr)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Warnw("failed to close grpc connection", "error", err)
		}
	}()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	client := v1.NewFiraServiceClient(conn)
	regErr := v1.RegisterFiraServiceHandlerClient(ctx, rmux, client)
	if regErr != nil {
		return fmt.Errorf("faild to register grpc client: %w", regErr)
	}

	a.mux.Handle("/api/v1/", rmux)

	return nil
}

func (a *App) StartHTTP(ctx context.Context) {
	log := a.Logger().Sugar().Named("startup")

	a.StartService(ctx, "http-server", func(ctx context.Context, errChan chan error) Finalizer {
		server := &http.Server{
			Addr:    a.cfg.BindHttp,
			Handler: a.mux,
		}

		go func() {
			defer func() {
				if p := recover(); p != nil {
					errChan <- fmt.Errorf("recovered from panic: %#v", p)
				}
			}()

			listener, listenErr := net.Listen("tcp", a.cfg.BindHttp)
			if listenErr != nil {
				errChan <- fmt.Errorf("failed to listen: %w", listenErr)
			}

			log.Infow("HTTP server ready", "addr", a.cfg.BindHttp)

			errChan <- server.Serve(listener)
		}()

		return server.Shutdown
	})

	//  handle the swagger UI
	a.mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/protos/fira/v1/api.swagger.json")
	})

	a.mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./dist/swagger-ui"))))
}

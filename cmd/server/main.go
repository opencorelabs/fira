package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/api"
	"github.com/opencorelabs/fira/internal/application"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, appErr := application.NewApp()
	if appErr != nil {
		fmt.Println("app init error:", appErr)
		os.Exit(1)
	}
	log := app.Logger().Named("startup").Sugar()

	startFrontendError := app.StartFrontend(ctx)
	if startFrontendError != nil {
		log.Fatalw("unable to start frontend server", "error", startFrontendError)
	}

	svc := &api.FiraApiService{}

	// start the gRPC server
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterFiraServiceServer(grpcServer, svc)
	log.Infow("gRPC server ready on localhost:5566...")
	go grpcServer.Serve(lis)

	// dial the gRPC server above to make a client connection
	conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	client := v1.NewFiraServiceClient(conn)
	err = v1.RegisterFiraServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}

	mux := app.Mux()

	// handle the API interface
	mux.Handle("/api/v1/", rmux)

	//  handle the swagger UI
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/protos/fira/v1/api.swagger.json")
	})

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./dist/swagger-ui"))))

	log.Infof("HTTP server ready on %s...", app.Config().BindHttp)

	err = http.ListenAndServe(app.Config().BindHttp, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvDefault(envName, defaultVal string) string {
	val, hasVal := os.LookupEnv(envName)
	if !hasVal {
		return defaultVal
	}
	return val
}

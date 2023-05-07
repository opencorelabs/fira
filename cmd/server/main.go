package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/api"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	svc := &api.FiraApiService{}

	// start the gRPC server
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterFiraServiceServer(grpcServer, svc)
	log.Println("gRPC server ready on localhost:5566...")
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

	// create a standard HTTP router
	mux := http.NewServeMux()

	// handle the API interface
	mux.Handle("/api/v1/", rmux)

	// handle the client
	go runFrontend()
	proxy, proxyErr := NewProxy("http://localhost:3000")
	if proxyErr != nil {
		log.Fatal("proxy error", proxyErr)
	}

	mux.HandleFunc("/", ProxyRequestHandler(proxy))

	//  handle the swagger UI
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/protos/fira/v1/api.swagger.json")
	})

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./dist/swagger-ui"))))

	bindTo := getEnvDefault("BIND", "localhost:8080")
	log.Printf("HTTP server ready on %s...\n", bindTo)

	err = http.ListenAndServe(bindTo, mux)
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

func runFrontend() {
	cmd := exec.Command("yarn", "run", "start")
	cmd.Dir = "./client"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("error starting client: %s\n", err)
	}
}

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

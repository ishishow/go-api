package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"20dojo-online/pkg/interfaces/interceptor"
	"20dojo-online/pkg/interfaces/rpc/collection"
	"20dojo-online/pkg/interfaces/rpc/gacha"
	"20dojo-online/pkg/interfaces/rpc/game"
	"20dojo-online/pkg/interfaces/rpc/ranking"
	"20dojo-online/pkg/interfaces/rpc/user"
)

var (
	gwPort     string
	serverPort string
	promAddr   string
)

func main() {
	gwPort = "8080"
	serverPort = "8081"
	promAddr = ":9100"
	runPrometheus()

	flag.Parse()
	defer glog.Flush()
	if err := run(serverPort, gwPort); err != nil {
		glog.Fatal(err)
	}
}

// runPrometheus runs prometheus metrics server. This is non-blocking function.
func runPrometheus() {
	mux := http.NewServeMux()
	// Enable histogram
	grpc_prometheus.EnableHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Prometheus metrics bind address", promAddr)
		log.Fatal(http.ListenAndServe(promAddr, mux))
	}()
}

func run(serverPort string, gwPort string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMetadata(requestIDAnnotator),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(":" + serverPort)
	if err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	if err := game.RegisterGameServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	if err := gacha.RegisterGachaServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	if err := ranking.RegisterRankingServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	if err := collection.RegisterCollectionServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Accept-Encoding", "Accept", "x-token"}),
	)(mux)

	log.Printf("gateway port:" + gwPort)
	log.Printf("server listen: " + serverPort)
	return http.ListenAndServe(":"+gwPort, handler)
}

// requestIDAnnotator takes requestID from http request header and sets it to metadata.
func requestIDAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	requestID := req.Header.Get(interceptor.XRequestIDKey)
	if requestID == "" {
		requestID = ""
	}

	return metadata.New(map[string]string{
		interceptor.XRequestIDKey: requestID,
	})
}

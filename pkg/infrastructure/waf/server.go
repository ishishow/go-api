package waf

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"20dojo-online/pkg/interfaces/interceptor"
	"20dojo-online/pkg/interfaces/registry"
	"20dojo-online/pkg/interfaces/rpc/collection"
	"20dojo-online/pkg/interfaces/rpc/gacha"
	"20dojo-online/pkg/interfaces/rpc/game"
	"20dojo-online/pkg/interfaces/rpc/ranking"
	"20dojo-online/pkg/interfaces/rpc/user"
)

func Serve(serverport string) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":"+serverport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userService := registry.InitializeUserService()
	gameService := registry.InitializeGameService()
	gachaService := registry.InitializeGachaService()
	rankingService := registry.InitializeRankingService()
	collectionService := registry.InitializeCollectionService()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.RequestIDInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)

	user.RegisterUserServiceServer(grpcServer, &userService)
	game.RegisterGameServiceServer(grpcServer, &gameService)
	gacha.RegisterGachaServiceServer(grpcServer, &gachaService)
	ranking.RegisterRankingServiceServer(grpcServer, &rankingService)
	collection.RegisterCollectionServiceServer(grpcServer, &collectionService)
	grpc_prometheus.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

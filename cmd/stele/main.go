package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dollarkillerx/stele/internal/config"
	"github.com/dollarkillerx/stele/internal/server"
	"github.com/dollarkillerx/stele/rpc/generate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.InitConfig()
	if cfg.Debug {
		log.Println(cfg)
		log.SetFlags(log.Llongfile | log.LstdFlags)
	}

	stele := server.NewStele(cfg)
	log.Println("Stele Run In: ", cfg.SocketAddr)
	listen, err := net.Listen("tcp", cfg.SocketAddr)
	if err != nil {
		log.Fatalln(err)
	}

	creds, err := credentials.NewServerTLSFromFile("cert/servermpe", "cert/serveryek")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer(
		grpc.Creds(creds),
		grpc.StreamInterceptor(streamInterceptor(cfg.Username, cfg.Password)),
		grpc.UnaryInterceptor(unaryInterceptor(cfg.Username, cfg.Password)),
	)

	generate.RegisterSteleServer(server, stele)
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

func streamInterceptor(username string, password string) func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := authorize(username, password)(stream.Context()); err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func unaryInterceptor(username string, password string) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := authorize(username, password)(ctx); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func authorize(username string, password string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if len(md["username"]) > 0 && md["username"][0] == username &&
				len(md["password"]) > 0 && md["password"][0] == password {
				return nil
			}

			return fmt.Errorf("AccessDeniedErr")
		}

		return fmt.Errorf("EmptyMetadataErr")
	}
}

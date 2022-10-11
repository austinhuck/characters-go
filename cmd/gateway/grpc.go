package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"sync"

	character "github.com/austinhuck/characters-go/gen/character"
	characterpb "github.com/austinhuck/characters-go/gen/grpc/character/pb"
	charactersvr "github.com/austinhuck/characters-go/gen/grpc/character/server"
	inventorypb "github.com/austinhuck/characters-go/gen/grpc/inventory/pb"
	inventorysvr "github.com/austinhuck/characters-go/gen/grpc/inventory/server"
	itempb "github.com/austinhuck/characters-go/gen/grpc/item/pb"
	itemsvr "github.com/austinhuck/characters-go/gen/grpc/item/server"
	inventory "github.com/austinhuck/characters-go/gen/inventory"
	item "github.com/austinhuck/characters-go/gen/item"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, characterEndpoints *character.Endpoints, itemEndpoints *item.Endpoints, inventoryEndpoints *inventory.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		characterServer *charactersvr.Server
		itemServer      *itemsvr.Server
		inventoryServer *inventorysvr.Server
	)
	{
		characterServer = charactersvr.New(characterEndpoints)
		itemServer = itemsvr.New(itemEndpoints)
		inventoryServer = inventorysvr.New(inventoryEndpoints)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpcmiddleware.WithUnaryServerChain(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	characterpb.RegisterCharacterServer(srv, characterServer)
	itempb.RegisterItemServer(srv, itemServer)
	inventorypb.RegisterInventoryServer(srv, inventoryServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Printf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Printf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Printf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}

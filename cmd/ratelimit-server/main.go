// Program ratelimit-server is gRPC server that throttles all API requests.
package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/marselester/apigate/internal/pb"
)

func main() {
	grpcAddr := flag.String("grpc", ":5000", "gRPC listen address")
	flag.Parse()

	// gRPC server for throttling API requests.
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("could not listen to gRPC port: %v", err)
	}
	grpcserver := grpc.NewServer()
	pb.RegisterRateLimitServiceServer(
		grpcserver,
		&server{},
	)
	// gRPC reflection provides information about publicly-accessible gRPC services on a server,
	// and assists clients (e.g., grpcurl) at runtime to construct RPC requests and responses
	// without precompiled service information. Read more about reflection at
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md.
	reflection.Register(grpcserver)

	log.Fatal(grpcserver.Serve(grpcListener))
}

// server is gRPC server that implements RateLimitServiceServer interface.
// It's like HTTP multiplexer.
type server struct{}

// ShouldRateLimit must respond to the request with an OK or OVER_LIMIT code.
func (s *server) ShouldRateLimit(ctx context.Context, req *pb.RateLimitRequest) (*pb.RateLimitResponse, error) {
	// Descriptors is a list of labels on which the rate limit service can base
	// its decision to accept or reject the request.
	log.Printf("%s: %+v", req.GetDomain(), req.GetDescriptors())

	resp := pb.RateLimitResponse{
		OverallCode: pb.RateLimitResponse_OVER_LIMIT,
	}
	return &resp, nil
}

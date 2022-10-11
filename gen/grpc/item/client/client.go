// Code generated by goa v3.9.1, DO NOT EDIT.
//
// item gRPC client
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package client

import (
	itempb "github.com/austinhuck/characters-go/gen/grpc/item/pb"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli itempb.ItemClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the item service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: itempb.NewItemClient(cc),
		opts:    opts,
	}
}
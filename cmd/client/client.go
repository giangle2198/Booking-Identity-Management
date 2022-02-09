package client

import (
	"context"
	"time"

	clientpb "booking-identity-management/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type (
	cimClient struct {
		clientpb.BIMAPIClient
	}
)

func NewBIMClient(address string) (clientpb.BIMAPIClient, error) {

	conn, err := grpc.DialContext(context.Background(),
		address, grpc.WithInsecure(),
		grpc.WithTimeout(2*time.Second), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		return nil, err
	}
	return &cimClient{BIMAPIClient: clientpb.NewBIMAPIClient(conn)}, nil
}

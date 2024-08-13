package stigg

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	SidecarServiceClient

	conn *grpc.ClientConn
}

func NewClient(addr string) (*Client, error) {
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	client := Client{
		SidecarServiceClient: NewSidecarServiceClient(conn),
		conn:                 conn,
	}

	return &client, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

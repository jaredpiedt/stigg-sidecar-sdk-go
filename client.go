package stigg

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Client manages communication with the Sidecar API.
type Client struct {
	SidecarServiceClient

	conn *grpc.ClientConn
}

// Newclient should be used to initialize a new Sidecar API client.
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

// Close closes the connection to the Sidecar API.
func (c *Client) Close() error {
	return c.conn.Close()
}

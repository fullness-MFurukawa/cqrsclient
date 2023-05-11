package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CommandConnector struct {
}

func (c *CommandConnector) Connect() (*grpc.ClientConn, error) {
	server := "commandservice:8082"
	conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

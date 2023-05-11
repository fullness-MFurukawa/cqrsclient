package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type QueryConnector struct {
}

/*
QueryServiceとの接続を確立しする
*/
func (c *QueryConnector) Connect() (*grpc.ClientConn, error) {
	server := "queryservice:8083"
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

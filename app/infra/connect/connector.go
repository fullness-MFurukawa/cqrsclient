package connect

import "google.golang.org/grpc"

// サーバへの接続インターフェイス
type ServerConnector interface {
	// サーバへ接続する
	Connect() (*grpc.ClientConn, error)
}

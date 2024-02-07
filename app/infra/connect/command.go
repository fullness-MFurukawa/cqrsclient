package connect

import (
	"crypto/tls"
	"crypto/x509"
	"embed"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//go:embed commandservice.pem
var commandFile embed.FS

// CommandServerと接続する
type commandConnector struct {
}

// コンストラクタ
func NewcommandConnector() ServerConnector {
	return &commandConnector{}
}

// CommandServiceとの接続を確立する
func (c *commandConnector) Connect() (*grpc.ClientConn, error) {
	// 埋め込まれた証明書を読み込む
	pemData, err := commandFile.ReadFile("commandservice.pem")
	if err != nil {
		return nil, err
	}
	// 証明書のプールを作成し、PEMデータから証明書を追加する
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(pemData) {
		return nil, err
	}
	config := &tls.Config{ // クライアント用のTLS設定
		RootCAs: cp,
	}
	creds := credentials.NewTLS(config)
	server := "commandservice:8082" // サーバ名とポート番号
	if conn, err := grpc.Dial(      // gRPCサーバとの接続を確立する
		server,                               // 接続文字列
		grpc.WithTransportCredentials(creds), // TLSを使用する
	); err != nil {
		return nil, err
	} else {
		return conn, nil
	}
}

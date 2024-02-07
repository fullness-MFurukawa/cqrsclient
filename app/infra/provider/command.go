package provider

import (
	"cqrsclient/infra/connect"

	"github.com/fullness-MFurukawa/samplepb/pb"
)

// Commandクライアントプロバイダ
type CommandClientProvider struct {
	Category pb.CategoryCommandClient
	Product  pb.ProductCommandClient
}

// コンストラクタ
func NewCommandClientProvider(connector connect.ServerConnector) (*CommandClientProvider, error) {
	// Command Serviceの接続
	if connect, err := connector.Connect(); err != nil {
		return nil, err
	} else {
		// カテゴリ用クライアントを生成する
		category := pb.NewCategoryCommandClient(connect)
		// 商品用クライアントを生成する
		product := pb.NewProductCommandClient(connect)
		// プロバイダの生成
		return &CommandClientProvider{Category: category, Product: product}, nil
	}
}

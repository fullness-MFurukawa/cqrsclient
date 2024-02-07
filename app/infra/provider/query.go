package provider

import (
	"cqrsclient/infra/connect"

	"github.com/fullness-MFurukawa/samplepb/pb"
)

// Queryクライアントプロバイダ
type QueryClientProvider struct {
	Category pb.CategoryQueryClient // カテゴリ用クライアント
	Product  pb.ProductQueryClient  // 商品用クライアント
}

// コンストラクタ
func NewQueryClientProvider(connector connect.ServerConnector) (*QueryClientProvider, error) {
	if connect, err := connector.Connect(); err != nil {
		return nil, err
	} else {
		// カテゴリ用クライアントを生成する
		category := pb.NewCategoryQueryClient(connect)
		// 商品用クライアントを生成する
		product := pb.NewProductQueryClient(connect)
		// プロバイダの生成
		return &QueryClientProvider{Category: category, Product: product}, nil
	}
}

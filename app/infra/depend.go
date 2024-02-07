package infra

import (
	"cqrsclient/infra/connect"
	"cqrsclient/infra/provider"

	"go.uber.org/fx"
)

// インフラストラクチャ層の依存定義
var InfraDepend = fx.Options(
	fx.Provide(
		// 生成されたインスタンスに名称CommandConnectorを付ける
		fx.Annotated{Name: "CommandConnector", Target: connect.NewcommandConnector},
		// 生成されたインスタンスに名称QueryConnectorを付ける
		fx.Annotated{Name: "QueryConnector", Target: connect.NewqueryConnector},
		fx.Annotated{
			Target: func(
				in struct {
					fx.In // 構造体を埋め込む
					// 名称CommandConnectorのインスタンスを設定する
					Connector connect.ServerConnector `name:"CommandConnector"`
				}) (*provider.CommandClientProvider, error) {
				// CommandClientProviderのコンストラクタ
				return provider.NewCommandClientProvider(in.Connector)
			},
		},
		fx.Annotated{
			Target: func(
				in struct {
					fx.In // 構造体を埋め込む
					// 名称QueryConnectorのインスタンスを設定する
					Connector connect.ServerConnector `name:"QueryConnector"`
				}) (*provider.QueryClientProvider, error) {
				// QueryClientProviderのコンストラクタ
				return provider.NewQueryClientProvider(in.Connector)
			},
		},
	),
)

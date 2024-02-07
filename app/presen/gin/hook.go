package gin

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

// fxのライフサイクル
func RegisterHooks(lifecycle fx.Lifecycle, router *Router) {
	lifecycle.Append(
		fx.Hook{
			// fx開始時の処理
			OnStart: func(context.Context) error {
				fmt.Println("CQRSClientの開始 Post:8081 !!!")
				go router.Engine.Run(":8081") //	Ginの起動
				return nil
			},
			// fx終了時の処理
			OnStop: func(context.Context) error {
				fmt.Println("CQRSClientの停止 !!!")
				return nil
			},
		},
	)
}

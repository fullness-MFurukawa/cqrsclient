package main

import (
	"client/presentation/gin"
	"fmt"
)

// @title gRPCサンプルプログラム(CQRSサービス)
// @version 1.0
// @description 商品カテゴリ情報と商品情報を管理するAPIサービス
// @termOfService http://localhost:8081/

// @contact.name 株式会社フルネス
// @contact.url https://www.fullness.co.jp
// @contact.email furukawa@fullness.co.jp

// @licence.name Apache 2.0
// @licence.url https://www.apache.org/licenses/LICENSE-2.0

// @host localhost:8081
// @Basepath /
func main() {
	fmt.Println("Gin Server起動!!!")
	// Ginの起動
	gin.RunServer()
}

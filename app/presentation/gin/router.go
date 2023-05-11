package gin

import (
	_ "client/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() {

	router := gin.Default()
	// CROS設定
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	// リクエストハンドラの生成
	query_handler, err := NewQueryHandler()
	if err != nil {
		panic(err)
	}
	command_handler, err := NewCommandHandler()
	if err != nil {
		panic(err)
	}

	// Swaggerの設定
	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// リクエストハンドラの登録(Query)
	router.GET("/category/list", query_handler.CategoryList)
	router.GET("/product/list", query_handler.ProductList)
	router.GET("/category/byid/:id", query_handler.CategoryById)
	router.GET("/product/byid/:id", query_handler.ProductById)
	router.GET("/product/bykeyword/:keyword", query_handler.ProductByKeyword)
	// リクエストハンドラの登録(Command)
	router.POST("/category/add", command_handler.CreateCategory)
	router.PUT("/category/update", command_handler.UpdateCategory)
	router.DELETE("/category/delete", command_handler.DeleteCategory)
	router.POST("/product/add", command_handler.CreateProduct)
	router.PUT("/product/update", command_handler.UpdateProduct)
	router.DELETE("/product/delete", command_handler.DeleteProduct)
	// Ginの起動
	router.Run(":8081")
}

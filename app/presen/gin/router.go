package gin

import (
	_ "cqrsclient/docs"
	"cqrsclient/presen/gin/handler"

	"github.com/gin-gonic/gin"
	sfiles "github.com/swaggo/files"
	gswag "github.com/swaggo/gin-swagger"
)

// ルータ構造体
type Router struct {
	Engine *gin.Engine
}

// コンストラクタ
func NewRouter(command_handler *handler.CommandHandler, query_handler *handler.QueryHandler) *Router {
	gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode) // リリースモードに設定する
	engine := gin.Default()      // デフォルトエンジンを生成する

	// Swaggerの設定
	url := gswag.URL("http://localhost:8081/swagger/doc.json")
	engine.GET("/swagger/*any", gswag.WrapHandler(sfiles.Handler, url))

	// リクエストハンドラの登録(Query)
	engine.GET("/category/list", query_handler.CategoryList)
	engine.GET("/product/list", query_handler.ProductList)
	engine.GET("/category/byid/:id", query_handler.CategoryById)
	engine.GET("/product/byid/:id", query_handler.ProductById)
	engine.GET("/product/bykeyword/:keyword", query_handler.ProductByKeyword)
	engine.GET("/product/liststream", query_handler.ProductListStream)
	// リクエストハンドラの登録(Command)
	engine.POST("/category/add", command_handler.CreateCategory)
	engine.PUT("/category/update", command_handler.UpdateCategory)
	engine.DELETE("/category/delete", command_handler.DeleteCategory)
	engine.POST("/product/add", command_handler.CreateProduct)
	engine.PUT("/product/update", command_handler.UpdateProduct)
	engine.DELETE("/product/delete", command_handler.DeleteProduct)

	router := Router{Engine: engine}
	return &router
}

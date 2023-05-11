package gin

import (
	"client/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryHandler struct {
	service application.QueryService
}

// すべての商品カテゴリを取得する
// @Summary 商品カテゴリを取得する
// @Info 2023/05/10
// @Description すべての商品カテゴリを取得する
// @ID get-categories
// @Accept */*
// @Produce json
// @Success 200 {object} pb.Category
// @Failure 500 {object} pb.Error
// @Router /category/list [GET]
func (h *QueryHandler) CategoryList(ctx *gin.Context) {
	result, err := h.service.CagtegoryList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.GetError() != nil {
		ctx.JSON(http.StatusInternalServerError, result.GetError())
		return
	}
	ctx.JSON(http.StatusOK, result.GetCategories())
}

// すべての商品を取得する
// @Summary 商品を取得する
// @Info 2023/05/10
// @Description すべての商品を取得する
// @ID get-products
// @Accept */*
// @Produce json
// @Success 200 {object} pb.Product
// @Failure 404 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/list [GET]
func (h *QueryHandler) ProductList(ctx *gin.Context) {
	result, err := h.service.ProductList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if result.GetError() != nil {
		ctx.JSON(http.StatusInternalServerError, result.GetError())
		return
	}
	ctx.JSON(http.StatusOK, result.GetProducts())
}

// 商品カテゴリを取得する
// @Summary 商品カテゴリを取得する
// @Info 2023/05/10
// @Description 指定されたカテゴリ番号の商品カテゴリを取得する
// @ID get-category-by-id
// @Accept */*
// @Produce json
// @Param id path string true "id(カテゴリ番号)"
// @Success 200 {object} pb.Category
// @Failure 404 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/byid/{id} [GET]
func (h *QueryHandler) CategoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("id = ", id)
	result, err := h.service.SearchCategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.GetError() != nil {
		ctx.JSON(http.StatusNotFound, result.GetError())
		return
	}
	ctx.JSON(http.StatusOK, result.GetCategory())
}

// 商品を取得する
// @Summary 商品を取得する
// @Info 2023/05/10
// @Description 指定された商品番号の商品を取得する
// @ID get-product-by-id
// @Accept */*
// @Produce json
// @Param id path string true "id(商品番号)"
// @Success 200 {object} pb.Product
// @Failure 404 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/byid/{id} [GET]
func (h *QueryHandler) ProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.service.SearchProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if result.GetError() != nil {
		ctx.JSON(http.StatusNotFound, result.GetError())
		return
	}
	ctx.JSON(http.StatusOK, result.GetProduct())
}

// 商品を取得する
// @Summary 商品を取得する
// @Info 2023/05/10
// @Description 指定されたキーワードが含まれる商品を取得する
// @ID get-product-by-keyword
// @Accept */*
// @Produce json
// @Param keyword path string true "キーワード"
// @Success 200 {object} pb.Product
// @Failure 404 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/bykeyword/{keyword} [GET]
func (h *QueryHandler) ProductByKeyword(ctx *gin.Context) {
	keyword := ctx.Param("keyword")
	result, err := h.service.SearchProductByKeyword(keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if result.GetError() != nil {
		ctx.JSON(http.StatusNotFound, result.GetError())
		return
	}
	ctx.JSON(http.StatusOK, result.GetProducts())
}

func NewQueryHandler() (*QueryHandler, error) {
	service, err := application.NewQueryService()
	if err != nil {
		return nil, err
	}
	return &QueryHandler{service: *service}, nil
}

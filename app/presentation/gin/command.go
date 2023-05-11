package gin

import (
	"client/application"
	"client/presentation/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommandHandler struct {
	service application.CommandService
}

// 商品カテゴリ登録を提供するハンドラ
// @Summary 商品カテゴリを登録する
// @Info 2023/05/10
// @Description 商品カテゴリを登録する
// @ID add-category
// @Accept application/json
// @Produce json
// @Param newcategory body dto.CategoryDTO true "name(カテゴリ名)"
// @Success 200 {object} pb.CategoryUpResult
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/add [post]
func (c *CommandHandler) CreateCategory(ctx *gin.Context) {
	var dto dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateCategory(dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 商品カテゴリ変更を提供するハンドラ
// @Summary 商品カテゴリを変更する
// @Info 2023/05/10
// @Description 商品カテゴリを変更する
// @ID update-category
// @Accept application/json
// @Produce json
// @Param newcategory body dto.CategoryDTO true "カテゴリ番号,カテゴリ名"
// @Success 200 {object} pb.CategoryUpResult
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/update [put]
func (c *CommandHandler) UpdateCategory(ctx *gin.Context) {
	var dto dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.UpdateCategory(dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 商品カテゴリ削除を提供するハンドラ
// @Summary 商品カテゴリを削除する
// @Info 2023/05/10
// @Description 商品カテゴリを削除する
// @ID delete-category
// @Accept application/json
// @Produce json
// @Param newcategory body dto.CategoryDTO true "カテゴリ番号,カテゴリ名"
// @Success 200 {object} pb.CategoryUpResult
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/delete [delete]
func (c *CommandHandler) DeleteCategory(ctx *gin.Context) {
	var dto dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.DeleteCategory(dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 商品登録を提供するハンドラ
// @Summary 商品を登録する
// @Info 2023/05/10
// @Description 商品を登録する
// @ID add-product
// @Accept application/json
// @Produce json
// @Param newproduct body dto.ProductDTO true "name(商品名),price(単価),categoryId(商品カテゴリ番号)"
// @Success 200 {object} pb.ProductUpResult
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/add [post]
func (c *CommandHandler) CreateProduct(ctx *gin.Context) {
	var dto dto.ProductDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateProduct(dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 商品変更を提供するハンドラ
// @Summary 商品を変更する
// @Info 2023/05/10
// @Description 商品を変更する
// @ID update-product
// @Accept application/json
// @Produce json
// @Param updateproduct body dto.ProductDTO true "id(商品番号),name(商品名),price(単価),categoryId(商品カテゴリ番号)"
// @Success 200 {object} pb.ProductUpResult
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/update [put]
func (c *CommandHandler) UpdateProduct(ctx *gin.Context) {
	var dto dto.ProductDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.UpdateProduct(dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// 商品削除を提供するハンドラ
// @Summary 商品を削除する
// @Info 2023/05/10
// @Description 商品を削除する
// @ID delete-product
// @Accept application/json
// @Produce json
// @Param deleteproduct body dto.ProductDTO true "商品番号"
// @Success 200 {object} pb.ProductUpResult
// @Failure 400 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/delete [delete]
func (c *CommandHandler) DeleteProduct(ctx *gin.Context) {
	var dto dto.ProductDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.DeleteProduct(dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func NewCommandHandler() (*CommandHandler, error) {
	service, err := application.NewCommandService()
	if err != nil {
		return nil, err
	}
	return &CommandHandler{service: *service}, nil
}

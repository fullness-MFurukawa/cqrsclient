package handler

import (
	"cqrsclient/presen/gin/helper"

	"github.com/gin-gonic/gin"
)

// 問合せリクエストハンドラ
type QueryHandler struct {
	helper *helper.QueryHelper
}

// コンストラクタ
func NewQueryHandler(helper *helper.QueryHelper) *QueryHandler {
	return &QueryHandler{helper: helper}
}

// カテゴリ一覧を取得する
// @tags カテゴリ
// @Summary 一覧を取得する
// @Description カテゴリ一覧を取得する
// @ID get-categories
// @Accept */*
// @Produce json
// @Success 200 {object} pb.Category
// @Failure 500 {object} pb.Error
// @Router /category/list [GET]
func (ins *QueryHandler) CategoryList(ctx *gin.Context) {
	ins.helper.CategoryList(ctx)
}

// 商品一覧を取得する
// @tags 商品
// @Summary 一覧を取得する(Server streaming RPC)
// @Description 商品一覧を取得する
// @ID get-productstream
// @Accept */*
// @Produce json
// @Success 200 {object} []pb.Product
// @Failure 500 {object} pb.Error
// @Router /product/liststream [GET]
func (ins *QueryHandler) ProductListStream(ctx *gin.Context) {
	ins.helper.ProductList(ctx)
}

// 商品一覧を取得する
// @tags 商品
// @Summary 一覧を取得する
// @Description 商品一覧を取得する
// @ID get-products
// @Accept */*
// @Produce json
// @Success 200 {object} pb.Product
// @Failure 500 {object} pb.Error
// @Router /product/list [GET]
func (ins *QueryHandler) ProductList(ctx *gin.Context) {
	ins.helper.ProductList(ctx)
}

// カテゴリを取得する
// @tags カテゴリ
// @Summary 取得する
// @Description 指定されたカテゴリIDのカテゴリを取得する
// @ID get-category-by-id
// @Accept */*
// @Produce json
// @Param id path string true "id(カテゴリID)"
// @Success 200 {object} pb.Category
// @Failure 404 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /category/byid/{id} [GET]
func (ins *QueryHandler) CategoryById(ctx *gin.Context) {
	ins.helper.CategoryById(ctx)
}

// 商品を取得する
// @tags 商品
// @Summary 取得する
// @Description 指定された商品IDの商品を取得する
// @ID get-product-by-id
// @Accept */*
// @Produce json
// @Param id path string true "id(商品ID)"
// @Success 200 {object} pb.Product
// @Failure 404 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/byid/{id} [GET]
func (ins *QueryHandler) ProductById(ctx *gin.Context) {
	ins.helper.ProductById(ctx)
}

// 商品を取得する
// @tags 商品
// @Summary 取得する
// @Description 指定されたキーワードが含まれる商品を取得する
// @ID get-product-by-keyword
// @Accept */*
// @Produce json
// @Param keyword path string true "キーワード"
// @Success 200 {object} pb.Product
// @Failure 404 {object} pb.Error
// @Failure 500 {object} pb.Error
// @Router /product/bykeyword/{keyword} [GET]
func (ins *QueryHandler) ProductByKeyword(ctx *gin.Context) {
	ins.helper.ProductByKeyword(ctx)
}

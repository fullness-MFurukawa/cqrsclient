package helper

import (
	"cqrsclient/infra/provider"
	"net/http"

	"github.com/fullness-MFurukawa/samplepb/pb"
	"github.com/gin-gonic/gin"
)

// 更新ヘルパ
type CommandHelper struct {
	provider *provider.CommandClientProvider
}

// コンストラクタ
func NewCommandHelper(provider *provider.CommandClientProvider) *CommandHelper {
	return &CommandHelper{provider: provider}
}

// カテゴリ登録をする
func (ins *CommandHelper) CreateCategory(ctx *gin.Context) {
	var param pb.CategoryUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := ins.provider.Category.Create(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

// 商品カテゴリを変更する
func (ins *CommandHelper) UpdateCategory(ctx *gin.Context) {
	var param pb.CategoryUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := ins.provider.Category.Update(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

// カテゴリを削除する
func (ins *CommandHelper) DeleteCategory(ctx *gin.Context) {
	var param pb.CategoryUpParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	if result, err := ins.provider.Category.Delete(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.Error)
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

// 商品を登録する
func (ins *CommandHelper) CreateProduct(ctx *gin.Context) {
	var param pb.ProductUpParam
	// JSON形式のリクエストパラメータをpb.ProductParamにマッピングする
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	// gRPCクライアント経由でCommand ServiceのCreate()メソッドを実行する
	if result, err := ins.provider.Product.Create(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil { // エラーになった場合
			ErrResp(ctx, result.Error) // エラーレスポンスを返す
			return
		}
		// エラーでなければ、pb.ProductのJSONをレスポンスとして返す
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}

// 商品を変更する
func (ins *CommandHelper) UpdateProduct(ctx *gin.Context) {
	var param pb.ProductUpParam
	// JSON形式のリクエストパラメータをpb.ProductParamにマッピングする
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	// gRPCクライアント経由でCommand ServiceのUpdate()メソッドを実行する
	if result, err := ins.provider.Product.Update(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil { // エラーになった場合
			ErrResp(ctx, result.Error) // エラーレスポンスを返す
			return
		}
		// エラーでなければ、pb.ProductのJSONをレスポンスとして返す
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}

// 商品を削除する
func (ins *CommandHelper) DeleteProduct(ctx *gin.Context) {
	var param pb.ProductUpParam
	// JSON形式のリクエストパラメータをpb.ProductParamにマッピングする
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrResp(ctx, err)
		return
	}
	// gRPCクライアント経由でCommand ServiceのDelete()メソッドを実行する
	if result, err := ins.provider.Product.Delete(ctx, &param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil { // エラーになった場合
			ErrResp(ctx, result.Error) // エラーレスポンスを返す
			return
		}
		// エラーでなければ、pb.ProductのJSONをレスポンスとして返す
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}

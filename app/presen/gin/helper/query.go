package helper

import (
	"cqrsclient/infra/provider"
	"io"
	"net/http"

	"github.com/fullness-MFurukawa/samplepb/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

// 問合せヘルパ
type QueryHelper struct {
	provider *provider.QueryClientProvider
}

// コンストラクタ
func NewQueryHelper(provider *provider.QueryClientProvider) *QueryHelper {
	return &QueryHelper{provider: provider}
}

// カテゴリ一覧を取得する
func (ins *QueryHelper) CategoryList(ctx *gin.Context) {
	param := &emptypb.Empty{} // 空パラメータの生成
	// サーバメソッドの呼び出し
	if result, err := ins.provider.Category.List(ctx, param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil { // エラーが返された?
			ErrResp(ctx, result.GetError())
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategories()) // OKレスポンスを送信
	}
}

// 商品一覧を取得する(Server streaming RPC)
func (ins *QueryHelper) ProductListStream(ctx *gin.Context) {
	param := &emptypb.Empty{} // 空パラメータを生成する
	// ListStream()メソッドを実行してストリームを受け取る
	if stream, err := ins.provider.Product.ListStream(ctx, param); err != nil {
		ErrResp(ctx, err) // エラーレスポンスを返す
	} else {
		var products = make([]*pb.Product, 0) // 受信した結果を格納するスライス
		for {
			product, err := stream.Recv() // ストリームからデータを受信する
			if err == io.EOF {            // EOFなら終了する
				break
			}
			if err != nil {
				ErrResp(ctx, err) // エラーレスポンスを返す
			}
			products = append(products, product) // スライスに受信結果を追加する
		}
		ctx.JSON(http.StatusOK, products)
	}
}

// 商品一覧を取得する
func (ins *QueryHelper) ProductList(ctx *gin.Context) {
	param := &emptypb.Empty{} // 空パラメータを生成する
	// gRPCクライアント経由でQuery ServiceのList()メソッドを実行する
	if result, err := ins.provider.Product.List(ctx, param); err != nil {
		ErrResp(ctx, err) // エラーレスポンスを返す
	} else {
		if result.GetError() != nil { // エラーになった場合
			ErrResp(ctx, result.GetError()) // エラーレスポンスを返す
			return
		}
		// エラーでなければ、[]pb.ProductのJSONをレスポンスとして返す
		ctx.JSON(http.StatusOK, result.GetProducts())
	}
}

// カテゴリを取得する
func (ins *QueryHelper) CategoryById(ctx *gin.Context) {
	param := &pb.CategoryParam{Id: ctx.Param("id")}
	if result, err := ins.provider.Category.ById(ctx, param); err != nil {
		ErrResp(ctx, err)
	} else {
		if result.GetError() != nil {
			ErrResp(ctx, result.GetError())
			return
		}
		ctx.JSON(http.StatusOK, result.GetCategory())
	}
}

// 商品を取得する
func (ins *QueryHelper) ProductById(ctx *gin.Context) {
	// URL中のパラメータを取得し、pb.ProductParamを生成する
	param := &pb.ProductParam{Id: ctx.Param("id")}
	// gRPCクライアント経由でQuery ServiceのById()メソッドを実行する
	if result, err := ins.provider.Product.ById(ctx, param); err != nil {
		ErrResp(ctx, err) // エラーレスポンスを返す
	} else {
		if result.GetError() != nil { // エラーになった場合
			ErrResp(ctx, result.GetError()) // エラーレスポンスを返す
			return
		}
		// エラーでなければ、pb.ProductのJSONをレスポンスとして返す
		ctx.JSON(http.StatusOK, result.GetProduct())
	}
}

// 商品を取得する
func (ins *QueryHelper) ProductByKeyword(ctx *gin.Context) {
	// URL中のパラメータを取得し、pb.ProductParamを生成する
	param := &pb.ProductParam{Keyword: ctx.Param("keyword")}
	// gRPCクライアント経由でQuery ServiceのBykeyword()メソッドを実行する
	if result, err := ins.provider.Product.ByKeyword(ctx, param); err != nil {
		ErrResp(ctx, err) // エラーレスポンスを返す
	} else {
		if result.GetError() != nil { // エラーになった場合
			ErrResp(ctx, result.GetError()) // エラーレスポンスを返す
			return
		}
		// エラーでなければ、[]pb.ProductのJSONをレスポンスとして返す
		ctx.JSON(http.StatusOK, result.GetProducts())
	}
}

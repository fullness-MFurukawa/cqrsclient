package helper

import (
	"net/http"

	"github.com/fullness-MFurukawa/samplepb/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// エラーレスポンス送信
func ErrResp(ctx *gin.Context, err any) {
	if pbErr, ok := err.(*pb.Error); ok {
		if pbErr.Type == "Domain Error" || pbErr.Type == "CRUD Error" {
			ctx.JSON(http.StatusBadRequest, pbErr)
		} else {
			ctx.JSON(http.StatusInternalServerError, pbErr)
		}
	} else {
		other, _ := err.(error)
		// エラーは*status.Statusに変換可能?
		if status, ok := status.FromError(other); ok {
			// ステータスはBad Request?
			if status.Code() == codes.InvalidArgument {
				ctx.JSON(http.StatusBadRequest,
					pb.Error{Type: "Validate Error", Message: status.Message()})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError,
				pb.Error{Type: "InterError Error", Message: other.Error()})
		}
	}
}

package v1

import (
	"github.com/gin-gonic/gin"
)

func HelloWorld() gin.HandlerFunc {
	//return func(ctx *gin.Context) {
	//	var err = errors.New("hello world")
	//	e.Fail(ctx, e.SERVE_INTERNAL.WithOrigin(err))
	//}

	//return func(ctx *gin.Context) {
	//	e.Fail(ctx, e.SERVE_INTERNAL.WithTips("hello world"))
	//}

	//return func(ctx *gin.Context) {
	//	e.Success(ctx, "hello world")
	//}

	return func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	}
}

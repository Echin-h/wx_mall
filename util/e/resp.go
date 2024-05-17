package e

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseBody struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, data ...any) {
	response := responseBody{
		Code: http.StatusOK,
		Msg:  SUCCESS.Message,
		Data: data,
	}
	//if len(data) > 0 {
	//	response.Data = data[0]
	//}
	c.JSON(http.StatusOK, response)
}

func Fail(c *gin.Context, err error) {
	var e *Error
	ok := errors.As(err, &e)
	if !ok {
		e = INTERNAL_SERVER_ERROR.WithOrigin(err)
	}

	var resp responseBody
	resp.Msg = e.Message
	resp.Data = e.Data
	resp.Code = e.Code

	c.JSON(int(e.Code/100), resp)
	c.Abort()
}

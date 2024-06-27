// 通用响应封装
package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
  Code int    `json:"code"`
  Data any    `json:"data"`
  Msg  string `json:"msg"`
}

const (
  Ok = 0
  Error   = 7
)


//通用的result
func Result(code int, data any, msg string, c *gin.Context) {
  c.JSON(http.StatusOK, Response{
    Code: code,
    Data: data,
    Msg:  msg,
  })
}

//请求成功的响应
func Success(data any, msg string, c *gin.Context) {
  Result(Ok, data, msg, c)
}
func SuccessWithData(data any, c *gin.Context) {
  Result(Ok, data, "成功", c)
}
func SuccessWithMessage(msg string, c *gin.Context) {
  Result(Ok, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
  Result(Error, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context) {
  Result(Error, map[string]any{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
  msg, ok := ErrorMap[code]
  if ok {
    Result(int(code), map[string]any{}, msg, c)
    return
  }
  Result(Error, map[string]any{}, "未知错误", c)
}

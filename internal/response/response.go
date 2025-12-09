package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 代表 handler 层的执行结果封装
type Response struct {
	// Code 状态码，非0时表示自定义HTTP业务状态码
	Code int `json:"code"`

	// Data 响应数据，具体内容因接口而异
	Data any `json:"data,omitempty"`

	// Msg 返回信息，通常是状态描述
	Msg string `json:"msg"`

	// Err 错误信息，序列化时忽略（仅供内部日志使用）
	Err error `json:"-"`
}

// Execute 包装器，自动根据 Response 返回统一格式的 HTTP 响应
func Execute(fn func(ctx *gin.Context) Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := fn(ctx)
		if res.Err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": res.Code,
				"msg":  res.Msg,
				"err":  res.Err.Error(),
			})
			return
		}

		// 支持自定义 code
		if res.Code != 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": res.Code,
				"data": res.Data,
				"msg":  res.Msg,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": res.Data,
				"msg":  res.Msg,
			})
		}
	}
}

// OK 成功响应
func OK(data any, msg string) Response {
	return Response{
		Code: 0,
		Data: data,
		Msg:  msg,
	}
}

// Fail 失败响应
func Fail(code int, msg string, err error) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

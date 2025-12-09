package handler

import (
	"github.com/gin-gonic/gin"
	commonModel "github.com/horonlee/ginhub/internal/model/common"
	res "github.com/horonlee/ginhub/internal/response"
	"github.com/horonlee/ginhub/internal/service"
)

type HelloWorldHandler struct {
	hwService service.HelloWorldService
}

func NewHelloWorldHandler(hwService service.HelloWorldService) *HelloWorldHandler {
	return &HelloWorldHandler{hwService: hwService}
}

// PostHelloWorld 处理POST /helloworld请求
func (h *HelloWorldHandler) PostHelloWorld() gin.HandlerFunc {
	return res.Execute(func(ctx *gin.Context) res.Response {
		// Handler层的请求体
		var req struct {
			Message string `json:"message" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			return res.Response{Msg: "Invalid request", Err: err}
		}

		// 调用Service层，只传递业务字段
		if err := h.hwService.PostHelloWorld(ctx.Request.Context(), req.Message); err != nil {
			return res.Response{Msg: "Failed to create hello world", Err: err}
		}

		// 获取数据库信息
		dbInfo, err := h.hwService.GetDatabaseInfo(ctx.Request.Context())
		if err != nil {
			return res.Response{Msg: "Failed to get database info", Err: err}
		}

		// Handler层的响应体
		return res.Response{
			Data: gin.H{
				"message":  req.Message,
				"version":  commonModel.Version,
				"database": dbInfo,
			},
			Msg: "success",
		}
	})
}

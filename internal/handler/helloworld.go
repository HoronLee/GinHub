package handler

import (
	commonModel "github.com/HoronLee/GinHub/internal/model/common"
	"github.com/HoronLee/GinHub/internal/model/helloworld"
	res "github.com/HoronLee/GinHub/internal/response"
	"github.com/HoronLee/GinHub/internal/service"
	"github.com/gin-gonic/gin"
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
		var req helloworld.CreateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			return res.Response{Msg: "", Err: err} // Msg 为空，自动使用 err 详情
		}

		if err := h.hwService.PostHelloWorld(ctx.Request.Context(), req.Message); err != nil {
			return res.Response{Msg: "Failed to create hello world", Err: err}
		}

		dbInfo, err := h.hwService.GetDatabaseInfo(ctx.Request.Context())
		if err != nil {
			return res.Response{Msg: "Failed to get database info", Err: err}
		}

		return res.Response{
			Data: helloworld.CreateResponse{
				Message:  req.Message,
				Version:  commonModel.Version,
				Database: dbInfo,
			},
			Msg: "success",
		}
	})
}

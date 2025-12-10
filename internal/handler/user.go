package handler

import (
	"github.com/HoronLee/GinHub/internal/model/user"
	res "github.com/HoronLee/GinHub/internal/response"
	"github.com/HoronLee/GinHub/internal/service"
	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建UserHandler实例
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register 用户注册处理器
func (h *UserHandler) Register() gin.HandlerFunc {
	return res.Execute(func(ctx *gin.Context) res.Response {
		var req user.RegisterRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			return res.Response{Msg: "Invalid request body", Err: err}
		}

		if err := h.userService.Register(ctx.Request.Context(), req); err != nil {
			return res.Response{Msg: "Registration failed", Err: err}
		}

		return res.Response{
			Data: gin.H{"message": "User registered successfully"},
			Msg:  "success",
		}
	})
}

// Login 用户登录处理器
func (h *UserHandler) Login() gin.HandlerFunc {
	return res.Execute(func(ctx *gin.Context) res.Response {
		var req user.LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			return res.Response{Msg: "Invalid request body", Err: err}
		}

		token, err := h.userService.Login(ctx.Request.Context(), req)
		if err != nil {
			return res.Response{Msg: "Login failed", Err: err}
		}

		return res.Response{
			Data: user.LoginResponse{Token: token},
			Msg:  "success",
		}
	})
}

// DeleteUser 删除用户处理器
func (h *UserHandler) DeleteUser() gin.HandlerFunc {
	return res.Execute(func(ctx *gin.Context) res.Response {
		// 从JWT中间件获取用户ID（当前登录用户）
		userIDValue, exists := ctx.Get("user_id")
		if !exists {
			return res.Response{Msg: "User not authenticated", Err: nil}
		}

		userID, ok := userIDValue.(uint)
		if !ok {
			return res.Response{Msg: "Invalid user ID format", Err: nil}
		}

		// 也可以从URL参数获取要删除的用户ID（如果需要管理员删除其他用户）
		// 这里简化为删除当前登录用户
		if err := h.userService.DeleteUser(ctx.Request.Context(), userID); err != nil {
			return res.Response{Msg: "Failed to delete user", Err: err}
		}

		return res.Response{
			Data: gin.H{"message": "User deleted successfully"},
			Msg:  "success",
		}
	})
}

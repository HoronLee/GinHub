package helloworld

// CreateRequest 创建 HelloWorld 的请求
type CreateRequest struct {
	Message string `json:"message" binding:"required,min=1,max=20"`
}

// CreateResponse 创建 HelloWorld 的响应
type CreateResponse struct {
	Message  string `json:"message"`
	Version  string `json:"version"`
	Database string `json:"database"`
}

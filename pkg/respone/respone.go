package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JsonResponse
type JsonResponse struct {
	Code     int         `json:"code"`    // 状态码
	Message  string      `json:"message"` // 消息内容
	Data     interface{} `json:"data"`    // 返回结构数据
	HttpCode int         // http状态码
}

// ToJson 响应json
func (resp *JsonResponse) ToJson(ctx *gin.Context) {
	code := 200
	if resp.HttpCode != 200 {
		code = resp.HttpCode
	}
	ctx.JSON(code, resp)
}

// FailResponse 失败响应
func FailResponse(code int, message string, data ...interface{}) *JsonResponse {
	var r interface{}
	if len(data) > 0 {
		r = data
	} else {
		r = struct{}{}
	}
	return &JsonResponse{
		Code:    code,
		Message: message,
		Data:    r,
	}
}

func SuccessResponse(data ...interface{}) *JsonResponse {
	var r interface{}
	if len(data) > 0 {
		r = data[0]
	} else {
		r = struct{}{}
	}
	return &JsonResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    r,
	}
}

func ErrorResponse(status int, message string, data ...interface{}) *JsonResponse {
	var r interface{}
	if len(data) > 0 {
		r = data
	} else {
		r = struct{}{}
	}

	return &JsonResponse{
		Code:    status,
		Message: message,
		Data:    r,
	}
}

// WriteTo 将 json 设为响应体.
// HTTP 状态码由应用状态码决定
func (resp *JsonResponse) WriteTo(ctx *gin.Context) {
	var code int
	if resp.HttpCode == 0 {
		code = http.StatusOK
	} else {
		code = resp.HttpCode
	}

	ctx.JSON(code, resp)
}
func (resp *JsonResponse) SetHttpCode(httpCode int) *JsonResponse {
	resp.HttpCode = httpCode
	return resp
}

// responseCode 获取 HTTP 状态码. HTTP 状态码由 应用状态码映射
func (that *JsonResponse) responseCode() int {
	// todo 完善应用状态码对应 http 状态码
	if that.Code != http.StatusOK {
		return 200
	}
	return 200
}

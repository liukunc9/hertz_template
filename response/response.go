package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

const (
	ERROR   = 500
	SUCCESS = 0
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(c *app.RequestContext, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *app.RequestContext) {
	Result(c, SUCCESS, nil, "操作成功")
}

func OkWithMessage(c *app.RequestContext, message string) {
	Result(c, SUCCESS, nil, message)
}

func OkWithData(c *app.RequestContext, data interface{}) {
	Result(c, SUCCESS, data, "查询成功")
}

func OkWithDetailed(c *app.RequestContext, data interface{}, message string) {
	Result(c, SUCCESS, data, message)
}

func Fail(c *app.RequestContext) {
	Result(c, ERROR, map[string]interface{}{}, "操作失败")
}

func FailWithMessage(c *app.RequestContext, message string) {
	Result(c, ERROR, map[string]interface{}{}, message)
}

func FailWithDetailed(c *app.RequestContext, data interface{}, message string) {
	Result(c, ERROR, data, message)
}

package hengo

import (
	"encoding/json"
	"github.com/kelin428/hengo/binding"
	"net/http"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request

	path   string
	method string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w:      w,
		r:      r,
		path:   r.URL.String(),
		method: r.Method,
	}
}

// JsonResponse 返回json格式的响应
// 处理请求
func (c *Context) JsonResponse(code StatusResponseCode, msg string, hData H) {
	c.w.Header().Set("Content-Type", "application/json")

	response := Response{
		Code: code,
		Msg:  msg,
		Data: &hData,
	}

	responseJson, _ := json.Marshal(response)

	c.w.WriteHeader(int(code))
	c.w.Write(responseJson)
}

func (c *Context) JsonErrorResponse(code StatusResponseCode, msg string) {
	c.JsonResponse(code, msg, nil)
}

// ShouldBindWith 绑定请求体到对象
// 暂时只支持json
func (c *Context) ShouldBindWith(obj any, b binding.Binding) error {
	return b.Bind(c.r, obj)
}

// MustBindWith 绑定请求体到对象
// 暂时只支持json
func (c *Context) MustBindWith(obj any, b binding.Binding) error {
	if err := c.ShouldBindWith(obj, b); err != nil {
		c.JsonErrorResponse(StatusBadRequest, "参数错误")
		return err
	}
	return nil
}

func (c *Context) BindJSON(obj any) error {
	return c.MustBindWith(obj, binding.JSON)
}

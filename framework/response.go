package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// IResponse代表返回方法
type IResponse interface {
	// Json输出
	Json(obj interface{}) IResponse

	// Jsonp输出
	Jsonp(obj interface{}) IResponse

	//xml输出
	Xml(obj interface{}) IResponse

	// html输出
	Html(file string, obj interface{}) IResponse

	// string
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置200状态
	SetOkStatus() IResponse
}

func (ctx *Context) Jsonp(obj interface{}) IResponse {
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
	callback := template.JSEscapeString(callbackFunc)

	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		log.Println("response write fail, err :", err)
		return ctx
	}
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		log.Println("response write fail, err :", err)
		return ctx
	}
	ret, err := json.Marshal(obj)
	if err != nil {
		log.Println("obj marshal fail, err :", err)
		return ctx
	}
	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		log.Println("response write fail, err :", err)
		return ctx
	}
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		log.Println("response write fail, err :", err)
		return ctx
	}

	return ctx
}

func (ctx *Context) Xml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/html")
	ctx.responseWriter.Write(byt)
	return ctx
}

// html输出
func (ctx *Context) Html(file string, obj interface{}) IResponse {
	// 读取模版文件，创建template实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	// 执行Execute方法将obj和模版进行结合
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")
	return ctx
}

// string
func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

// 重定向
func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

// header
func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

// Cookie
func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

// 设置状态码
func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

// 设置200状态
func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}

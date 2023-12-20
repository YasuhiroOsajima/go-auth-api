package controllers

type Context interface {
	JSON(code int, obj any)
	Param(key string) string
	ShouldBindJSON(obj any) error
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
}

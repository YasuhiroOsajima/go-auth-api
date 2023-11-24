package controllers

type Context interface {
	JSON(code int, obj any)
	Param(key string) string
	ShouldBindJSON(obj any) error
}

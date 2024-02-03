package api

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	Json(int, interface{})
}

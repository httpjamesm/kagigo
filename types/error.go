package types

type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Ref  interface{} `json:"ref"`
}

type ErrorResponse struct {
	Meta  interface{} `json:"meta"`
	Data  interface{} `json:"data"`
	Error []Error     `json:"error"`
}

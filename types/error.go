package types

type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Ref  interface{} `json:"ref"`
}

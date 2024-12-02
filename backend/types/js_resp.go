package types

type JSResp struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type JsObject map[string]any

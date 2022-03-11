package api

type HealtcheckRequestBody struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

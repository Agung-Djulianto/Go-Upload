package model

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ResponseSuccess struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseFailed struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}

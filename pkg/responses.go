package pkg

import (
	"test-hiring/config"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data,omitempty"`
	Meta        interface{} `json:"meta,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Token       interface{} `json:"token,omitempty"`
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
	ctx.Abort()
	return
}

func Responses(code int, data *config.Result) *Response {
	var respone = Response{
		Code:   code,
		Status: getStatus(code),
	}

	if data.Message != nil {
		respone.Description = data.Message
	}
	if data.Data != nil {
		respone.Data = data.Data
	}
	if data.Meta != nil {
		respone.Meta = data.Meta
	}
	if data.Token != nil {
		respone.Token = data.Token
	}

	return &respone
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
		break
	case 201:
		desc = "Created"
		break
	case 203:
		desc = "Success Modified"
		break
	case 400:
		desc = "Bad Request"
		break
	case 401:
		desc = "Unauthorized"
		break
	case 403:
		desc = "Forbidden"
		break
	case 404:
		desc = "Not Found"
		break
	case 500:
		desc = "Internal Server Error"
		break
	case 501:
		desc = "Bad Gateway"
		break
	case 304:
		desc = "Not Modified"
		break
	default:
		desc = ""
	}

	return desc
}

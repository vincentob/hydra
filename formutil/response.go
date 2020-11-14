package formutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MsgOK  = "success"
	MsgErr = "error"
)

type APIResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Errors string      `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &APIResponse{
		Status: 0,
		Msg:    MsgOK,
		Data:   data,
	})
}

func ResponseErr(c *gin.Context, code int, err error) {
	c.JSON(code, &APIResponse{
		Status: 1,
		Msg:    MsgErr,
		Errors: err.Error(),
	})
}

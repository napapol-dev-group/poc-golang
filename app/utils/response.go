package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	StatusFail    string = "fail"
	StatusSuccess string = "success"
)

type baseSuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type baseFailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	MessageOk string = "OK"
)

// JSONSuccessResponse success response
func JSONSuccessResponse(c *gin.Context, data interface{}) {
	r := new(baseSuccessResponse)
	r.Status = StatusSuccess
	r.Message = MessageOk
	r.Data = data
	c.AbortWithStatusJSON(http.StatusOK, *r)
}

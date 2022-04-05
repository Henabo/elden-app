package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	SUCCESS = "ok"
	ERROR   = "error"
)

func Result(status string, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		status,
		data,
		message,
	})
}

func OK(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OKWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OKWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

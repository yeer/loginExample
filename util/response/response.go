package response

import (
	"loginExample/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

type Context struct {
	*gin.Context
}

func (c *Context) JsonSuccess(data interface{}) {
	c.Context.JSON(http.StatusOK, Response{
		Code:   constants.StatusOk,
		ErrMsg: "success",
		Data:   data,
	})
	c.Abort()
}

func (c *Context) JsonError(code int, data string) {
	var errMsg string
	if code == constants.StatusOk {
		errMsg = data
	} else {
		errMsg = constants.StatusText(code)
	}
	c.Context.JSON(http.StatusOK, Response{
		Code:   code,
		ErrMsg: errMsg,
		Data:   "",
	})
	c.Abort()
}
func (c *Context) Redirect(data string) {
	c.Context.Redirect(http.StatusMovedPermanently, data)
	return
}

func (c *Context) Redirect302(data string) {
	c.Context.Redirect(http.StatusFound, data)
	return
}

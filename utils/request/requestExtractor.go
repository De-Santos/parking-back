package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"parking-back/initializers"
	"parking-back/obj"
)

func BindValidQuery(c *gin.Context, o any) *obj.Error {
	err := c.BindQuery(o)
	if err != nil {
		return &obj.Error{Message: "Invalid query params", Err: err}
	}

	err = initializers.V.Struct(o)
	if err != nil {
		return &obj.Error{Message: "Invalid request query: " + fmt.Sprint(err), Err: err}
	}
	return nil
}

func BindValidBody(c *gin.Context, o any) *obj.Error {
	err := c.Bind(o)
	if err != nil {
		return &obj.Error{Message: "Failed to read dto", Err: err}

	}

	err = initializers.V.Struct(o)
	if err != nil {
		return &obj.Error{Message: "Invalid request body: " + fmt.Sprint(err), Err: err}
	}
	return nil
}

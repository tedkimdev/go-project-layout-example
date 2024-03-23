package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.POST(path, handler...)
}

func (n *Network) registerPUT(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.DELETE(path, handler...)
}

func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func (n *Network) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusBadRequest, result)
}

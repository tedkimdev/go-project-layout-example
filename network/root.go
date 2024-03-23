package network

import (
	"tedkimdev/go-project-layout-example/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engin *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engin:   gin.New(),
		service: service,
	}

	newUserRouter(r, service.User)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}

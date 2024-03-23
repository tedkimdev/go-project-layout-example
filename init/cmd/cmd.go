package cmd

import (
	"tedkimdev/go-project-layout-example/config"
	"tedkimdev/go-project-layout-example/network"
	"tedkimdev/go-project-layout-example/repository"
	"tedkimdev/go-project-layout-example/service"
)

type Cmd struct {
	config *config.Config

	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}

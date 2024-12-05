package network

import (
	"code/jaeger"

	"github.com/gin-gonic/gin"
)


type Network struct {
	client *jaeger.Client
	engine *gin.Engine
}

func NewNetwork(client *jaeger.Client) *Network {
	n := &Network {client: client, engine: gin.New()}

	newRouter(n)

	return n
}

func (n *Network) Start() {
	n.engine.Run(":8080")
}

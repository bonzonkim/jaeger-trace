package network

import "github.com/gin-gonic/gin"


type R int8


const (
	GET R = iota
	POST
	DELETE
	PUT
)

func (n *Network) Router(r R, path string, handler gin.HandlerFunc) {
	switch r {
		case GET:
			n.engine.GET(path, handler)
		case POST:
			n.engine.POST(path, handler)
		case DELETE:
			n.engine.DELETE(path, handler)
		case PUT:
			n.engine.PUT(path, handler)
	}
}

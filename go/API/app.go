package api

import (
	"code/API/network"
	"code/jaeger"
)

type App struct {
	client *jaeger.Client
}

func NewApp(serviceName string) {
	a := &App{}

    a.client = jaeger.NewClient(serviceName);
	n := network.NewNetwork(a.client)	
	n.Start()
}

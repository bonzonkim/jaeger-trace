package jaeger

import (
	"code/jaeger/provider"
	"io"
)

type Client struct {
	closer io.Closer
}

func NewClient(serviceName string) *Client {
	c := &Client{}
	var err error

	if c.closer, err = provider.NewProvider(serviceName); err != nil {
		panic(err)
	} else {
		return c
	}
}

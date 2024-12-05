package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func newRootSpan(name string, c *gin.Context) opentracing.Span {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan(name, ext.RPCServerOption(spanCtx))

	return sendSpan
}

func (s *Router) defaultHandler(c *gin.Context) {
	c.JSON(404, gin.H{
		"error": "This path is not yet implemented",
	})
}

func (s *Router) send(c *gin.Context) {
	sendSpan := newRootSpan("send_root_span", c)
	defer sendSpan.Finish()
	c.JSON(http.StatusOK, "Success")
}


func (s *Router) sendWithTag(c *gin.Context) {
	rootSpan := newRootSpan("send_root_span_with_tag", c)
	rootSpan.SetTag("tag_one", "one 입니다")
	rootSpan.SetTag("tag_two", "two 입니다")
	defer rootSpan.Finish()

	c.JSON(http.StatusOK, "Success")
}

func (s *Router) sendWithChild(c *gin.Context) {
	rootSpan := newRootSpan("send_root_span_with_child", c)
	defer rootSpan.Finish()
	childSpan := opentracing.GlobalTracer().StartSpan("child_span", opentracing.ChildOf(rootSpan.Context()))
	defer childSpan.Finish()
	c.JSON(http.StatusOK, c)
}

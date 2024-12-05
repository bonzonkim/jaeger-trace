package network

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func (s *Router) sendForPanic (c *gin.Context) {
	// 1
	tracer := opentracing.GlobalTracer()
	rootSpan := newRootSpan("root_span_for_error", c)
	defer rootSpan.Finish()

	fmt.Println("send for panic")
	fmt.Println("send Header", c.Request.Header)
	fmt.Println()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-for-error", nil)
	tracer.Inject(rootSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	// Inject 함으로써 아래의 차일드 스팬에 스팬 컨텍스트를 넘긴다.

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	c.JSON(http.StatusOK, gin.H{"message" : "root_span_for_error successfully done"})
}


func (s *Router) receiveForError (c *gin.Context) {
	// 1
	tracer := opentracing.GlobalTracer()

	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	childSpan := tracer.StartSpan("receive_for_error", opentracing.ChildOf(spanCtx) )
	childSpan.SetTag("error", true)
	childSpan.LogFields(log.String("event", "error"), log.String("message", "에러다!!"))
	fmt.Println("receive for error")

	defer childSpan.Finish()
	c.JSON(http.StatusOK, gin.H{"message": "receve_for_error successfully done"})
}

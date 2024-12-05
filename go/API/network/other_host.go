package network

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)
func (s *Router) sendWithOtherHost(c *gin.Context) {
	// 1
	tracer := opentracing.GlobalTracer()
	rootSpan := newRootSpan("other_host_root_span", c)
	defer rootSpan.Finish()

	fmt.Println("send with other host")
	fmt.Println("send Header", c.Request.Header)
	fmt.Println()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-from-other-host", nil)
	tracer.Inject(rootSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	// Inject 함으로써 아래의 차일드 스팬에 스팬 컨텍스트를 넘긴다.

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	c.JSON(http.StatusOK, "Success")
}

func (s *Router) receiveOne(c *gin.Context) {
	fmt.Println("/receive Header", c.Request.Header)
	tracer := opentracing.GlobalTracer()
	// 2
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	// other_host_root_span으로 넘겨받은 스팬컨텍스트를 그대로 넘겨받고 그 컨텍스트를 기준으로 새로운 차일드 스팬을 생성.
	childSpan := tracer.StartSpan("reveive_one_span", opentracing.ChildOf(spanCtx))

	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/receive-two-from-other-host", nil)
	tracer.Inject(childSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	defer childSpan.Finish()

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	c.JSON(http.StatusOK, gin.H{"message": "Span Receive One"})

}

func (s *Router) receiveTwo(c *gin.Context) {
	// 3
	fmt.Println("/receive Header", c.Request.Header)

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	childSpan := tracer.StartSpan("reveive_two_span", opentracing.ChildOf(spanCtx))

	defer childSpan.Finish()
	c.JSON(http.StatusOK, gin.H{"message": "Span Receive Two"})
}

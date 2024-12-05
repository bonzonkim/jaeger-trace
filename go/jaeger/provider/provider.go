package provider

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
)

func NewProvider(service string) (io.Closer, error) {
	setting := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1, // 위의 Type이 이 Param에 의해 false,true가 정해짐
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}
	// docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
	//if tracer, closer, err := setting.NewTracer(config.Logger(log.StdLogger)); err != nil {
	//	return nil, err
	//} else {
	//	opentracing.SetGlobalTracer(tracer)
	//	return closer, nil
	//}
	tracer, closer, err := setting.NewTracer(config.Logger(log.StdLogger))
	if err != nil {
		return nil, err
	}
	log.StdLogger.Infof("Tracer created for service %s", service)
	opentracing.SetGlobalTracer(tracer)
	return closer, nil
}

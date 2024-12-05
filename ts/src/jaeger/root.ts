import * as opentracing from 'opentracing'
import * as jaeger from 'jaeger-client'


const initTracer = (serviceName: string) => {
  const cfg: jaeger.TracingConfig = {
    serviceName: serviceName,
    sampler: {
      type: "const",
      param: 1,
    },
    reporter: {
      logSpans: true,
    }
  };

  const opt: jaeger.TracingOptions = {
    logger: {
      info: (msg: string) => {
        console.log("INFO: ", msg)
      },
      error: (msg: string) => {
        console.log("ERROR: ", msg)
      }
    }
  }
  const tracer = jaeger.initTracer(cfg, opt)
  opentracing.initGlobalTracer(tracer);
}

export default initTracer

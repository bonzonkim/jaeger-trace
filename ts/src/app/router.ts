import * as express from 'express'
import * as opentracing from 'opentracing';
import newRootSpan from './util'

const apiRouter = express.Router()

apiRouter.get("/hello", (req, res) => {
  const tracer: opentracing.Tracer = opentracing.globalTracer();
  const rootSpan = newRootSpan("hello-span", req);

  rootSpan.setTag("hello", "test hello tag value");
  rootSpan.setTag(opentracing.Tags.HTTP_METHOD, req.method);
  rootSpan.setTag(opentracing.Tags.HTTP_URL, req.originalUrl);
  rootSpan.log({ event: 'request_received', value: 'Procession /hello' });

  const carrier: Record<string, string> = {};
  tracer.inject(rootSpan.context(), opentracing.FORMAT_HTTP_HEADERS, carrier);

  rootSpan.log({ event: "response_sent" });
  rootSpan.finish();

  res.json({
    "message": "root span success",
    "tracingHeaders": carrier,
  });
});

apiRouter.get("/hello-with-child", (req, res) => {
  const tracer: opentracing.Tracer = opentracing.globalTracer();

  const rootSpan = newRootSpan("hello-with-child", req);
  
  const childSpan = tracer.startSpan("child-span", { childOf: rootSpan.context() });

  childSpan.setTag(opentracing.Tags.HTTP_METHOD, req.method);
  childSpan.setTag(opentracing.Tags.HTTP_URL, req.originalUrl);
  childSpan.log({ event: "child_procession", value: "Processing child span" });

  rootSpan.finish();
  childSpan.finish();

  res.json({
    "message": "child span success",
    "tracingHeaders": req.headers,
  });
});

export default apiRouter

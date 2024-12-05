import { opentracing } from "jaeger-client";
import * as express from 'express';

export default function newRootSpan(spanName: string, req: express.Request): opentracing.Span {
  const tracer: opentracing.Tracer = opentracing.globalTracer();
  const spanContext = tracer.extract(opentracing.FORMAT_HTTP_HEADERS, req.headers);
  const rootSpan = tracer.startSpan(spanName, { childOf: spanContext });

  return rootSpan;
}

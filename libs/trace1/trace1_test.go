package trace1

import (
	"context"
	"fmt"
	"testing"

	"go.opencensus.io/trace"
)

func Test_t1(t *testing.T) {
	pCtx, traceSpan := trace.StartSpan(context.Background(), "ttt")
	defer traceSpan.End()

	tid := traceSpan.SpanContext().TraceID.String()
	sid := traceSpan.SpanContext().SpanID.String()

	t1(pCtx, tid, sid)
	t2(pCtx, tid, sid)
}

func t1(ctx context.Context, tid, sid string) {
	_, traceSpan := trace.StartSpan(ctx, "DownloadFile")
	defer traceSpan.End()

	fmt.Println(traceSpan.SpanContext().TraceID.String() == tid)
	fmt.Println(traceSpan.SpanContext().SpanID.String() != sid)
}

func t2(ctx context.Context, tid, sid string) {
	traceSpan := trace.FromContext(ctx)
	fmt.Println(traceSpan.SpanContext().TraceID.String() == tid)
	fmt.Println(traceSpan.SpanContext().SpanID.String() == sid)
}

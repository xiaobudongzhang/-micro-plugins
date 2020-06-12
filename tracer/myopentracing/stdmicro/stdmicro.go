package stdmicro

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/metadata"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	statusCode "github.com/xiaobudongzhang/micro-plugins/breaker/http"
)

var sf = 100

func init() {
	rand.Seed(time.Now().Unix())
}

func SetSamplingFrequency(n int) {
	sf = n
}

func TracerWrapper(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md := make(map[string]string)
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		sp := opentracing.GlobalTracer().StartSpan(r.URL.Path, opentracing.ChildOf(spanCtx))

		defer sp.Finish()

		if err := opentracing.GlobalTracer().Inject(
			sp.Context(),
			opentracing.TextMap,
			opentracing.TextMapCarrier(md)); err != nil {
			log.Println(err)
		}

		ctx := context.TODO()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		ctx = metadata.NewContext(ctx, md)
		r = r.WithContext(ctx)

		sct := &statusCode.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
		n.ServeHTTP(sct.WrappedResponseWriter(), r)

		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.EscapedPath())
		ext.HTTPStatusCode.Set(sp, uint16(sct.Status))
		if sct.Status > http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		} else if rand.Intn(100) > sf {
			ext.SamplingPriority.Set(sp, 0)
		}
	})
}

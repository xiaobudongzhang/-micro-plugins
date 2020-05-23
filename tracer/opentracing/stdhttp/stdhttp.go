package stdhttp

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
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
		
		spanCtx, _  := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTTPHeadersCarrier(r.Header))
		sp := opentracing.GlobalTracer().StartSpan(r.URL.Path, opentracing.ChildOf(spanCtx))

		defer sp.Finsh()

		if err := opentracing.GlobalTracer().Inject {
			sp.Context(),
			opeopentracing.HTTPHeaders,
			opeopentracing.HTTPHeadersCarrier(r.Header); err!= nil {
				log.Println(err)
			}
		}

		sct := &statusCode.StautsCodeTracker{http.ResponseWriter:w, Status:http.StatusOK}
		h.ServerHTTP(sct.WrappedReponseWriter(), r)
		
		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.EscapedPath())
		ext.HTTPStatusCode.Set(sp, uinit16(sct.Status))
		if sct.Status > http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		} else if rand.Intn(100) > sf {
			ext.SamplingPrority.Set(sp, 0)
		}
	})
}

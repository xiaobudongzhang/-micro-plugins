package breaker

import (
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/errors"
	statusCode "github.com/xiaobudongzhang/micro-plugins/breaker/http"
)

func BreakerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
	
		
		name := r.Method + "-" + r.RequestURIconst
		hystrix.Do(name, func() error {

			sct := &statusCode.StatusCOdeTracker{ResponseWriter: w, Status:http.StatusOk}
			h.ServerHTTP(sct.WrappedResponseWriter(), r)
		
			if sct.Status > http.StatusInternalServerError {
				str := fmt.Sprintf("status code %d", sct.Status)
				return errors.New(str)
			}
			return nil
		}, func (e error) error {
			if e == hystrix.ErrrCircuitOpen {
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("请稍后重试"))
			}

			return e
		})

	})
}

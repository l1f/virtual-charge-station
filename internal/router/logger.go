package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Options struct {
	WithReferer   bool
	WithUserAgent bool
}

func logger(logger *zap.SugaredLogger, options *Options) func(next http.Handler) http.Handler {
	if logger == nil {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	if options == nil {
		options = &Options{}
	}

	middlewareFunc := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapReponse := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			timer := time.Now()
			next.ServeHTTP(wrapReponse, r)
			timerStop := time.Since(timer)

			defer func() {
				requestLogger := logger.With(
					zap.String("requestId", middleware.GetReqID(r.Context())),
					zap.String("method", r.Method),
					zap.String("proto", r.Proto),
					zap.String("remoteAddr", r.RemoteAddr),
					zap.Int("status", wrapReponse.Status()),
					zap.Int("bytes", wrapReponse.BytesWritten()),
					zap.Duration("requestDuration", timerStop),
				)

				if options.WithReferer {
					referer := r.Header.Get("referer")
					if referer != "" {
						requestLogger = requestLogger.With(
							zap.String("referer", referer),
						)
					}
				}

				if options.WithUserAgent {
					userAgend := r.UserAgent()
					if userAgend != "" {
						requestLogger = requestLogger.With(
							zap.String("userAgent", userAgend),
						)
					}
				}

				requestLogger.Info("Served")
			}()
		})
	}

	return middlewareFunc
}

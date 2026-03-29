package core_http_middleware

import (
	"context"
	"net/http"
	"time"

	core_logger "github.com/Fitray/Todo-list/internal/core/logger"
	core_http_responce "github.com/Fitray/Todo-list/internal/core/trasport/http/responce"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestID_Header = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestID_Header)
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set(requestID_Header, requestID)
			w.Header().Set(requestID_Header, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestID_Header)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responceHandler := core_http_responce.NewHTTPResponceHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responceHandler.PanicResponce(
						p, "got unexpected panic during handle HTTP",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_responce.NewResponceWritter(w)

			before := time.Now()

			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", time.Now().UTC()),
			)

			next.ServeHTTP(rw, r)

			latency := time.Since(before)
			log.Debug(
				"<<< done HTTP request",
				zap.Int("status-code", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", latency),
			)
		})
	}
}

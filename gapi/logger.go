package gapi

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func GrpcLogger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	startTime := time.Now()
	resp, err = handler(ctx, req)
	duration := time.Since(startTime)

	statusCode := codes.Unknown

	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := log.Info()

	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.Str("protocol", "grpc").Int("status_code", int(statusCode)).Str("status_text", statusCode.String()).Dur("duration", duration).Str("method", info.FullMethod).Msg("received a grpc request")

	return resp, nil
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		rec := &ResponseRecorder{w, http.StatusOK, nil}
		handler.ServeHTTP(rec, r)

		duration := time.Since(startTime)

		logger := log.Info()

		logger.Str("protocol", "http").Str("method", r.Method).Str("path", r.RequestURI).Int("status_code", rec.StatusCode).Dur("duration", duration).Msg("received an http request")

	})
}

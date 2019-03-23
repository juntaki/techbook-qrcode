package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"cloud.google.com/go/datastore"
	"github.com/juntaki/techbook-qrcode/src/lib/qrcode"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type myGorpTracer struct {
	logger *zap.Logger
}

func (t *myGorpTracer) Printf(format string, v ...interface{}) {
	t.logger.Info("gorp SQL Trace", zap.String("sql", fmt.Sprintf(format, v...)))
}

func main() {
	start(os.Getenv("QRCODE_ENV"))
}

func start(env string) {
	// Middleware
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	// zap logging
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync() // flushes buffer, if any

	// GRPC Handlers
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}

	// Datastore
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			grpc_zap.UnaryServerInterceptor(zapLogger),
		),
	)

	qs := InitializeQRCodeServiceServer(dsClient)
	qrcode.RegisterQRCodeServiceServer(s, qs)

	wrappedGrpc := http.StripPrefix("/grpc-web", grpcweb.WrapServer(s))
	for _, endPoint := range grpcweb.ListGRPCResources(s) {
		r.Mount("/grpc-web"+endPoint, wrappedGrpc)
	}

	ts := InitializeTechBookServer(dsClient)
	r.Mount("/code/", ts)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// gRPC server needs its recovery middleware, global chi's one is not working.
func recoveryFunc(p interface{}) error {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	log.Printf("panic recovered: %+v", string(buf))
	return status.Errorf(codes.Internal, "%s", p)
}

package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/cmd/blogsearch/cmd"
	"gitlab.com/zerok/zerokspot.com/pkg/otelhandler"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

var localZoneName string
var localZone *time.Location
var tracer trace.Tracer

var rootCmd = &cobra.Command{
	Use:           "blog",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		localZone, err = time.LoadLocation(localZoneName)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().StringVar(&localZoneName, "tz", "Europe/Vienna", "Timezone to be used for data-relevant processing")
	rootCmd.AddCommand(cmd.RootCmd)
	rootCmd.AddCommand(generateServeCmd())
	rootCmd.AddCommand(generateResizePhotosCmd())
}

func findParentTrace(ctx context.Context) context.Context {
	var traceParent string
	for _, v := range []string{"OVERRIDE_TRACEPARENT", "TRACEPARENT"} {
		traceParent = os.Getenv(v)
		if traceParent != "" {
			break
		}
	}
	if traceParent == "" {
		slog.InfoContext(ctx, "No parent trace provided")
		return ctx
	}
	slog.InfoContext(ctx, "Parent trace provided", slog.String("traceparent", traceParent))
	carrier := make(propagation.MapCarrier)
	carrier.Set("traceparent", traceParent)
	prop := otel.GetTextMapPropagator()
	ctx = prop.Extract(ctx, carrier)
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return ctx
	}
	return trace.ContextWithRemoteSpanContext(ctx, span.SpanContext())
}

func initOtel(ctx context.Context) *sdktrace.TracerProvider {
	var exporter sdktrace.SpanExporter
	var err error

	slog.InfoContext(ctx, "Configuring OTEL",
		slog.String("OTEL_EXPORTER_OTLP_ENDPOINT", os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
		slog.String("OTEL_EXPORTER_OTLP_PROTOCOL", os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL")),
	)

	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != "" {
		os.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")+"/v1/traces")
		os.Setenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL"))
		os.Setenv("OTEL_EXPORTER_OTLP_TRACES_HEADERS", os.Getenv("OTEL_EXPORTER_OTLP_HEADERS"))

		otlpClient := otlptracehttp.NewClient()
		exporter, err = otlptrace.New(ctx, otlpClient)
		slog.InfoContext(ctx, "Sending traces to remote endpoint")
	} else {
		exporter, err = stdouttrace.New(stdouttrace.WithWriter(os.Stderr))
		slog.InfoContext(ctx, "Sending traces to stderr")
	}
	if err != nil {
		slog.ErrorContext(ctx, "Failed to initialize trace exporter", slog.Any("err", err))
		os.Exit(1)
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("zerokspot-cli"),
		),
	)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate OTLP resource", slog.Any("err", err))
		os.Exit(1)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	tracer = tp.Tracer("zerokspot-cli")
	return tp
}

func main() {
	logger := slog.New(otelhandler.OTELHandler{Handler: slog.NewTextHandler(os.Stderr, nil)})
	slog.SetDefault(logger)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	tp := initOtel(ctx)
	defer func() {
		if err := tp.ForceFlush(context.Background()); err != nil {
			slog.ErrorContext(ctx, "Failed to flush tracer provider", slog.Any("err", err))
		}
	}()
	defer func() {
		slog.InfoContext(ctx, "Shutting down tracer provider")
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.ErrorContext(ctx, "Failed to shut down tracer provider", slog.Any("err", err))
		}
	}()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		slog.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}

func requireStringFlags(cmd *cobra.Command, flags ...string) error {
	for _, flag := range flags {
		val, err := cmd.Flags().GetString(flag)
		if err != nil || val == "" {
			return fmt.Errorf("required flag `%s` not set", flag)
		}
	}
	return nil
}

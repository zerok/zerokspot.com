package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/cmd/blogsearch/cmd"
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

var logger zerolog.Logger
var localZoneName string
var localZone *time.Location
var tracer trace.Tracer

var rootCmd = &cobra.Command{
	Use: "blog",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
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
	logger := zerolog.Ctx(ctx)
	var traceParent string
	for _, v := range []string{"OVERRIDE_TRACEPARENT", "TRACEPARENT"} {
		traceParent = os.Getenv(v)
		if traceParent != "" {
			break
		}
	}
	if traceParent == "" {
		logger.Info().Msg("No parent trace provided")
		return ctx
	}
	logger.Info().Msgf("Parent trace provided: %s", traceParent)
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

	logger.Info().
		Str("OTEL_EXPORTER_OTLP_ENDPOINT", os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")).
		Str("OTEL_EXPORTER_OTLP_PROTOCOL", os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL")).
		Msg("Configuring Otel")

	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != "" {
		os.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")+"/v1/traces")
		os.Setenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL"))
		os.Setenv("OTEL_EXPORTER_OTLP_TRACES_HEADERS", os.Getenv("OTEL_EXPORTER_OTLP_HEADERS"))

		otlpClient := otlptracehttp.NewClient()
		exporter, err = otlptrace.New(ctx, otlpClient)
		logger.Info().Msg("Sending traces to remote endpoint")
	} else {
		exporter, err = stdouttrace.New(stdouttrace.WithWriter(os.Stderr))
		logger.Info().Msg("Sending traces to stderr")
	}
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize trace exporter")
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("zerokspot-cli"),
		),
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to generate OLTP resource")
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
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	tp := initOtel(ctx)
	defer func() {
		if err := tp.ForceFlush(context.Background()); err != nil {
			logger.Error().Err(err).Msg("Failed to flush tracer provider")
		}
	}()
	defer func() {
		logger.Info().Msg("Shutting down tracer provider")
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error().Err(err).Msg("Failed to shut down tracer provider")
		}
	}()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		logger.Fatal().Msg(err.Error())
	}
}

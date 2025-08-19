package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/goflash/flash/v2"
	"github.com/goflash/flash/v2/ctx"
	mw "github.com/goflash/flash/v2/middleware"
	"go.uber.org/zap"
)

// slog -> zap handler that preserves attributes and groups.
type zapHandler struct {
	z      *zap.Logger
	attrs  []slog.Attr
	groups []string
}

func (h *zapHandler) Enabled(_ context.Context, _ slog.Level) bool { return true }

func (h *zapHandler) Handle(_ context.Context, r slog.Record) error {
	fields := make([]zap.Field, 0, len(h.attrs)+r.NumAttrs())

	// base attrs from WithAttrs
	for _, a := range h.attrs {
		h.appendAttr(&fields, a, h.groups)
	}
	// record attrs
	r.Attrs(func(a slog.Attr) bool {
		h.appendAttr(&fields, a, h.groups)
		return true
	})

	logger := h.z.With(fields...)
	switch r.Level {
	case slog.LevelDebug:
		logger.Debug(r.Message)
	case slog.LevelInfo:
		logger.Info(r.Message)
	case slog.LevelWarn:
		logger.Warn(r.Message)
	case slog.LevelError:
		logger.Error(r.Message)
	default:
		logger.Info(r.Message)
	}
	return nil
}

func (h *zapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	merged := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	merged = append(merged, h.attrs...)
	merged = append(merged, attrs...)
	return &zapHandler{
		z:      h.z,
		attrs:  merged,
		groups: append([]string{}, h.groups...),
	}
}

func (h *zapHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return &zapHandler{
		z:      h.z,
		attrs:  append([]slog.Attr{}, h.attrs...),
		groups: append(append([]string{}, h.groups...), name),
	}
}

func (h *zapHandler) appendAttr(fields *[]zap.Field, a slog.Attr, groups []string) {
	// resolve LogValuer if present
	v := a.Value
	if v.Kind() == slog.KindAny {
		if lv, ok := v.Any().(slog.LogValuer); ok {
			v = lv.LogValue()
		}
	}

	prefix := strings.Join(groups, ".")
	key := a.Key
	if prefix != "" {
		key = prefix + "." + key
	}

	switch v.Kind() {
	case slog.KindString:
		*fields = append(*fields, zap.String(key, v.String()))
	case slog.KindInt64:
		*fields = append(*fields, zap.Int64(key, v.Int64()))
	case slog.KindUint64:
		*fields = append(*fields, zap.Uint64(key, v.Uint64()))
	case slog.KindFloat64:
		*fields = append(*fields, zap.Float64(key, v.Float64()))
	case slog.KindBool:
		*fields = append(*fields, zap.Bool(key, v.Bool()))
	case slog.KindDuration:
		*fields = append(*fields, zap.Duration(key, v.Duration()))
	case slog.KindTime:
		*fields = append(*fields, zap.Time(key, v.Time()))
	case slog.KindGroup:
		for _, ga := range v.Group() {
			h.appendAttr(fields, slog.Attr{Key: ga.Key, Value: ga.Value}, append(groups, a.Key))
		}
	default:
		*fields = append(*fields, zap.Any(key, v.Any()))
	}
}

func main() {
	zl, err := zap.NewProduction() // JSON, includes ts and level
	if err != nil {
		log.Fatal(err)
	}
	defer zl.Sync()

	app := flash.New()

	adapter := &zapHandler{z: zl}
	app.SetLogger(slog.New(adapter))

	app.Use(mw.Logger(), mw.Recover(), mw.RequestID())

	// Example of using the custom zap logger with slog adapter
	app.GET("/", func(c flash.Ctx) error {
		return c.String(http.StatusOK, "zap via slog adapter")
	})

	// Example of logging with custom attributes
	app.GET("/custom", func(c flash.Ctx) error {
		logger := ctx.LoggerFromContext(c.Context())
		logger.Info(
			"custom attributes log",
			slog.String("user_id", "12345"),
			slog.String("role", "admin"),
			slog.Group("request",
				slog.String("ip", c.Request().RemoteAddr),
				slog.String("agent", c.Request().UserAgent()),
			),
		)

		// Using groups and attributes
		logger.WithGroup("request").Info("logged with custom attributes")

		// Using WithAttrs to add custom attributes
		logger.With(slog.String("user_id", "12345")).Info("logged with custom attributes")

		return c.String(http.StatusOK, "logged with custom attributes")
	})

	log.Fatal(http.ListenAndServe(":8080", app))
}

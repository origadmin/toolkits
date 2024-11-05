package loge

import (
	"log/slog"

	"github.com/golang-cz/devslog"
	"github.com/lmittmann/tint"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	KindAny       = slog.KindAny
	KindBool      = slog.KindBool
	KindDuration  = slog.KindDuration
	KindFloat64   = slog.KindFloat64
	KindGroup     = slog.KindGroup
	KindInt64     = slog.KindInt64
	KindLogValuer = slog.KindLogValuer
	KindString    = slog.KindString
	KindTime      = slog.KindTime
	KindUint64    = slog.KindUint64
	LevelDebug    = slog.LevelDebug
	LevelError    = slog.LevelError
	LevelInfo     = slog.LevelInfo
	LevelKey      = slog.LevelKey
	LevelWarn     = slog.LevelWarn
	MessageKey    = slog.MessageKey
	SourceKey     = slog.SourceKey
	TimeKey       = slog.TimeKey

	Black        = devslog.Black
	Blue         = devslog.Blue
	Cyan         = devslog.Cyan
	Green        = devslog.Green
	Magenta      = devslog.Magenta
	Red          = devslog.Red
	UnknownColor = devslog.UnknownColor
	White        = devslog.White
	Yellow       = devslog.Yellow
)

type (
	Attr           = slog.Attr
	Handler        = slog.Handler
	HandlerOptions = slog.HandlerOptions
	JSONHandler    = slog.JSONHandler
	Kind           = slog.Kind
	Level          = slog.Level
	Leveler        = slog.Leveler
	LevelVar       = slog.LevelVar
	Logger         = slog.Logger
	LogValuer      = slog.LogValuer
	Record         = slog.Record
	Source         = slog.Source
	TextHandler    = slog.TextHandler
	Value          = slog.Value

	TintOptions = tint.Options

	Color          = devslog.Color
	DevslogOptions = devslog.Options

	LumberjackLogger = lumberjack.Logger
)

var (
	Any           = slog.Any
	AnyValue      = slog.AnyValue
	Bool          = slog.Bool
	BoolValue     = slog.BoolValue
	Debug         = slog.Debug
	DebugContext  = slog.DebugContext
	Default       = slog.Default
	Duration      = slog.Duration
	DurationValue = slog.DurationValue
	Error         = slog.Error
	ErrorContext  = slog.ErrorContext
	Float64       = slog.Float64
	Float64Value  = slog.Float64Value
	Group         = slog.Group
	GroupValue    = slog.GroupValue
	Info          = slog.Info
	InfoContext   = slog.InfoContext
	Int           = slog.Int
	Int64         = slog.Int64
	Int64Value    = slog.Int64Value
	IntValue      = slog.IntValue
	Log           = slog.Log
	LogAttrs      = slog.LogAttrs

	NewJSONHandler    = slog.NewJSONHandler
	NewLogLogger      = slog.NewLogLogger
	NewRecord         = slog.NewRecord
	NewTextHandler    = slog.NewTextHandler
	SetDefault        = slog.SetDefault
	SetLogLoggerLevel = slog.SetLogLoggerLevel
	String            = slog.String
	StringValue       = slog.StringValue
	Time              = slog.Time
	TimeValue         = slog.TimeValue
	Uint64            = slog.Uint64
	Uint64Value       = slog.Uint64Value
	Warn              = slog.Warn
	WarnContext       = slog.WarnContext
	With              = slog.With
	Err               = tint.Err
	NewTintHandler    = tint.NewHandler
	NewDevslogHandler = devslog.NewHandler
)

package simplelog

// LevelType ...
type LevelType int32

const (
	// NOLEVEL ...
	NOLEVEL LevelType = iota
	// DEBUG ...
	DEBUG
	// INFO ...
	INFO
	// WARN ...
	WARN
	// ERROR ...
	ERROR
	// PANIC ...
	PANIC
	// FATAL ...
	FATAL
)

// Level field name.
var (
	// NoLevelName ...
	NoLevelName = ""
	// DebugLevelName ...
	DebugLevelName = "debug"
	// InfoLevelName ...
	InfoLevelName = "info"
	// WarnLevelName ...
	WarnLevelName = "warn"
	// ErrorLevelName ...
	ErrorLevelName = "error"
	// PanicName ...
	PanicLevelName = "panic"
	// FatalLevelName ...
	FatalLevelName = "fatal"
)

func (level LevelType) String() string {
	switch level {
	case DEBUG:
		return DebugLevelName
	case INFO:
		return InfoLevelName
	case WARN:
		return WarnLevelName
	case ERROR:
		return ErrorLevelName
	case PANIC:
		return PanicLevelName
	case FATAL:
		return FatalLevelName
	}
	return NoLevelName
}

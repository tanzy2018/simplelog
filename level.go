package simplelog

// LevelType ...
type LevelType int32

const (
	_ LevelType = iota
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
	// NOLEVEL ...
	NOLEVEL
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

var levelNames = [...]string{
	"",
	DebugLevelName,
	InfoLevelName,
	WarnLevelName,
	ErrFieldName,
	PanicLevelName,
	FatalLevelName,
	NoLevelName,
}

func (level LevelType) String() string {
	if !level.isValid() {
		return ""
	}
	return levelNames[level]
}

func (level LevelType) isValid() bool {
	if level < DEBUG || level > NOLEVEL {
		return false
	}
	return true
}

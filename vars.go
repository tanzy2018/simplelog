package simplelog

import (
	"fmt"
	"time"

	"github.com/tanzy2018/simplelog/encode"
	"github.com/tanzy2018/simplelog/internal"
)

const (
	leftDelimiter  = '{'
	rightDelimiter = '}'
	endDelimiter   = '\n'
	fieldDelimiter = ','
	kvDelimiter    = ':'
	valueWrapper   = '"'
)

const (
	// TimestampUnixFormat ...
	TimestampUnixFormat = "unix"
	// TimestampUnixMilliFormat ...
	TimestampUnixMilliFormat = "unixmilli"
	// TimestampUnixMicroFormat ...
	TimestampUnixMicroFormat = "unixmicro"
	// TimestampUnixNanoFormat ...
	TimestampUnixNanoFormat = "unixnano"
	// .createTime_lastmodifiedTime
	renameFormat = ".%v_%v"
)

var (
	// EnableTimeField ...
	EnableTimeField = true
	// TimeFieldName ...
	TimeFieldName = "time"
	// TimeFieldFormat ...
	TimeFieldFormat = "2006-01-02 15:04:05"
	// LevelFieldName ...
	LevelFieldName = "level"
	// MsgFieldName ...
	MsgFieldName = "msg"
	// StackFieldName ...
	StackFieldName = "stack"
)

func timeMeta() encode.Meta {
	if TimeFieldFormat == TimestampUnixFormat {
		return encode.Int64(TimeFieldName, time.Now().Unix())
	}

	if TimeFieldFormat == TimestampUnixMilliFormat {
		return encode.Int64(TimeFieldName, time.Now().UnixNano()/1000000)
	}

	if TimeFieldFormat == TimestampUnixMicroFormat {
		return encode.Int64(TimeFieldName, time.Now().UnixNano()/1000)
	}

	if TimeFieldFormat == TimestampUnixNanoFormat {
		return encode.Int64(TimeFieldName, time.Now().UnixNano())
	}

	return encode.String(TimeFieldName, internal.TimeFormat(TimeFieldFormat))
}

func levelMeta(level LevelType) encode.Meta {
	return encode.String(LevelFieldName, level.String())
}

func msgMeta(msg string) encode.Meta {
	return encode.String(MsgFieldName, msg)
}

func stackMeta() encode.Meta {
	return encode.String(StackFieldName, internal.CallStack(6))
}

func genRenameSubfix(csec, msec int64) string {

	var v0, v1 interface{}
	t0, t1 := time.Unix(csec, 0), time.Unix(msec, 0)
	switch TimeFieldFormat {
	case TimestampUnixFormat, TimestampUnixMilliFormat, TimestampUnixMicroFormat, TimestampUnixNanoFormat:
		v0, v1 = t0.Unix(), t1.Unix()
	default:
		v0, v1 = t0.Format(TimeFieldFormat), t1.Format(TimeFieldFormat)

	}
	return fmt.Sprintf(renameFormat, v0, v1)
}

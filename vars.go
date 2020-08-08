package simplelog

import (
	"fmt"
	"time"

	"github.com/tanzy2018/simplelog/meta"
	"github.com/tanzy2018/simplelog/utils"
)

const (
	leftDelimiter  = '{'
	rightDelimiter = '}'
	endDelimiter   = '\n'
	fieldDelimiter = ','
	kvDelimiter    = ':'
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
	// UseTimeField ...
	UseTimeField = true
	// TimeFieldName ...
	TimeFieldName string = "time"
	// TimeFieldFormat ...
	TimeFieldFormat string = "2006-01-02 15:04:05"
	// LevelFieldName ...
	LevelFieldName string = "level"
	// MsgFieldName ...
	MsgFieldName string = "msg"
)

func timeMeta() meta.Meta {
	if TimeFieldFormat == TimestampUnixFormat {
		return meta.Int64(TimeFieldName, time.Now().Unix())
	}

	if TimeFieldFormat == TimestampUnixMilliFormat {
		return meta.Int64(TimeFieldName, time.Now().UnixNano()/1000000)
	}

	if TimeFieldFormat == TimestampUnixMicroFormat {
		return meta.Int64(TimeFieldName, time.Now().UnixNano()/1000)
	}

	if TimeFieldFormat == TimestampUnixNanoFormat {
		return meta.Int64(TimeFieldName, time.Now().UnixNano())
	}

	return meta.String(TimeFieldName, utils.TimeFormat(TimeFieldFormat))
}

func levelMeta(level LevelType) meta.Meta {
	return meta.String(LevelFieldName, level.String())
}

func msgMeta(msg string) meta.Meta {
	return meta.String(MsgFieldName, msg)
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

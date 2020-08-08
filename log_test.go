package simplelog

import (
	"io/ioutil"
	"os"

	// "runtime"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/tanzy2018/simplelog/meta"
	"github.com/tanzy2018/simplelog/utils"
)

var once sync.Once

func TestDefault_overall(t *testing.T) {
	defer Sync()
	AddHooks(func() meta.Meta {
		return meta.Int("socore", utils.RandInt(100))
	})

	Hook(func() meta.Meta {
		return meta.String("service", "demo")
	})

	Hook(func() meta.Meta {
		return meta.Int64("service-id", int64(utils.RandInt(1000)+1000))
	})
	Debug("debugmsg", meta.Int("uid", 11))
	Info("infomsg", meta.Int("uid", 12), meta.String("detail", "xxxxinfo...."))
	Warn("warnmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	Error("errmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	Panic("panicmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	// Fatal("fatalmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
}

func TestNew_overall(t *testing.T) {
	newLog := New()
	defer newLog.Sync()
	AddHooks(func() meta.Meta {
		return meta.Int("socore", utils.RandInt(100))
	})

	newLog.Hook(func() meta.Meta {
		return meta.String("service", "demo")
	})

	newLog.Hook(func() meta.Meta {
		return meta.Int64("service-id", int64(utils.RandInt(1000)+1000))
	})

	newLog.Debug("你好中国", meta.Int("uid", 11))
	newLog.Info("infomsg", meta.Int("uid", 12), meta.String("detail", "xxxxinfo...."))
	newLog.Warn("warnmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	newLog.Error("errmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	newLog.Panic("panicmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	// newLog.Fatal("fatalmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
}

func BenchmarkLog_default(b *testing.B) {
	TimeFieldFormat = TimestampUnixMilliFormat
	defer Sync()
	DeFault().WithWriterCloser(Discard, false, false)
	once.Do(func() {

		AddHooks(func() meta.Meta {
			return meta.Int("score", utils.RandInt(100))
		})

		Hook(func() meta.Meta {
			return meta.String("service", "demo")
		})

		Hook(func() meta.Meta {
			return meta.Int64("service-id", int64(utils.RandInt(1000)+1000))
		})

	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Info("infomsg", meta.Int("uid", 12), meta.String("detail", "xxxxinfo...."))
	}

}

func BenchmarkLog_new(b *testing.B) {
	// runtime.GOMAXPROCS(2)
	var newLog *Log
	newLog = New(
		WithMaxRecordSize(1024*10),
		WithMaxSyncSize(1024*1024),
		WithMaxFileSize(1024*1024*1024)).
		WithWriterCloser(Discard, false, true)
	// WithFileWriter("testdata", "", "overall-newsimple.txt")
	once.Do(func() {

		TimeFieldFormat = time.StampMilli
		score := utils.RandInt(100)
		AddHooks(func() meta.Meta {
			return meta.Int("score", score)
		})

		newLog.Hook(func() meta.Meta {
			return meta.String("service", "demo")
		})

		newLog.Hook(func() meta.Meta {
			return meta.String("from", "demo-service")
		})

		serverID := int64(utils.RandInt(1000) + 1000)
		newLog.Hook(func() meta.Meta {
			return meta.Int64("service-id", serverID)
		})

		randomStr := utils.RandomString(1024)
		newLog.Hook(func() meta.Meta {
			return meta.String("randomstr", randomStr)
		})
	})
	defer newLog.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newLog.Info("infomsg", meta.Int("uid", 12), meta.String("detail", "xxxxinfo...."))
	}

}

func TestZeroLog(t *testing.T) {
	zerolog.TimeFieldFormat = time.StampMilli
	zerolog.MessageFieldName = "msg"
	path := "testdata/overall.txt"
	file := openFile(path)
	defer file.Close()
	logger := zerolog.New(file).With().Timestamp().Logger()
	logger.Info().Str("detail", "你好中国").Int("uid", 12).Msg("infomsg")
}

func BenchmarkZeroLog(b *testing.B) {
	zerolog.TimeFieldFormat = time.StampMilli
	zerolog.MessageFieldName = "msg"
	score := utils.RandInt(100)
	serverID := int64(utils.RandInt(1000) + 1000)
	randomStr := utils.RandomString(1024)
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("detail", "xxxxinfo....").
			Int("uid", 12).
			Int("score", score).
			Str("service", "demo").
			Int64("service-id", serverID).
			Str("randomestr", randomStr).
			Str("from", "demo-service").
			Msg("infomsg")
	}
}

func openFile(path string) *os.File {
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	return file
}

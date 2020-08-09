package simplelog

import (
	"io/ioutil"
	"os"

	// "runtime"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/tanzy2018/simplelog/encode"
	"github.com/tanzy2018/simplelog/internal"
)

var once sync.Once

func TestSimpleLog(t *testing.T) {
	newLog := New(
		WithRecordBufsLen(20),
		WithSyncBufsLen(20),
	)
	defer newLog.Sync()
	AddHooks(func() encode.Meta {
		return encode.Int("socore", internal.RandInt(100))
	})

	newLog.Hook(func() encode.Meta {
		return encode.String("service", "demo")
	})

	newLog.Hook(func() encode.Meta {
		return encode.Int64("service-id", int64(internal.RandInt(1000)+1000))
	})

	time.Sleep(time.Millisecond * 500)

	newLog.Debug("你好中国", encode.Int("uid", 11))
	newLog.Info("infomsg", encode.Int("uid", 12), encode.String("detail", "xxxxinfo...."))
	newLog.Warn("warnmsg", encode.Int("uid", 13), encode.String("detail", "xxxxwarn...."))
	newLog.Error("errmsg", encode.Int("uid", 13), encode.String("detail", "xxxxwarn...."))
	newLog.Panic("panicmsg", encode.Int("uid", 13), encode.String("detail", "xxxxwarn...."))
	// newLog.Fatal("fatalmsg", encode.Int("uid", 13), encode.String("detail", "xxxxwarn...."))
}

func BenchmarkSimpleLog(b *testing.B) {
	// runtime.GOMAXPROCS(1)
	var newLog *Log
	newLog = New(
		WithRecordBufsLen(20),
		WithSyncBufsLen(20),
		WithMaxRecordSize(1024*10),
		WithMaxSyncSize(1024*1024),
		WithMaxFileSize(1024*1024*1024)).
		WithWriterCloser(Discard, false, true)
		// WithFileWriter("testdata", "", "simplelog.txt")
	once.Do(func() {

		TimeFieldFormat = time.StampMilli
		score := internal.RandInt(100)
		AddHooks(func() encode.Meta {
			return encode.Int("score", score)
		})

		newLog.Hook(func() encode.Meta {
			return encode.String("service", "demo")
		})

		newLog.Hook(func() encode.Meta {
			return encode.String("from", "demo-service")
		})

		serverID := int64(internal.RandInt(1000) + 1000)
		newLog.Hook(func() encode.Meta {
			return encode.Int64("service-id", serverID)
		})

		randomStr := internal.RandomString(1024)
		newLog.Hook(func() encode.Meta {
			return encode.String("randomstr", randomStr)
		})
	})
	defer newLog.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newLog.Info("infomsg", encode.Int("uid", 12), encode.String("detail", "xxxxinfo...."))
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
	score := internal.RandInt(100)
	serverID := int64(internal.RandInt(1000) + 1000)
	randomStr := internal.RandomString(1024)
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

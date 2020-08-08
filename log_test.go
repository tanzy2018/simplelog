package simplelog

import (
	"fmt"
	"log"
	"os"
	// "runtime"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"github.com/tanzy2018/simplelog/meta"
	"github.com/tanzy2018/simplelog/utils"
)

var once sync.Once

func TestDefault_overall(t *testing.T) {
	defer Sync()
	err := InitFileWriter("testdata",
		"", "overall-default.txt")
	if err != nil {
		panic(fmt.Errorf("panic open file:%v", err))
	}
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
	newLog, err := New(WithFileWriter(
		"testdata",
		"",
		"overall-new.txt"))
	if err != nil {
		panic(fmt.Errorf("panic open file:%v", err))
	}

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

	newLog.Debug("debugmsg", meta.Int("uid", 11))
	newLog.Info("infomsg", meta.Int("uid", 12), meta.String("detail", "xxxxinfo...."))
	newLog.Warn("warnmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	newLog.Error("errmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	newLog.Panic("panicmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
	// newLog.Fatal("fatalmsg", meta.Int("uid", 13), meta.String("detail", "xxxxwarn...."))
}

func BenchmarkLog_default(b *testing.B) {
	InitFileWriter("testdata",
		"", "overall-default.txt")
	SetDirectWrite(false)
	TimeFieldFormat = TimestampUnixMilliFormat
	defer Sync()
	once.Do(func() {

		AddHooks(func() meta.Meta {
			return meta.Int("socore", utils.RandInt(100))
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
	newLog, err := New(
		WithFileWriter(
			"testdata",
			"",
			"overall-new.txt"),
		WithMaxRecordSize(1024*10),
		WithMaxSyncSize(1024*1024),
		WithMaxFileSize(1024*1024*100))
	if err != nil {
		panic(fmt.Errorf("panic open file:%v", err))
	}

	defer newLog.Sync()

	once.Do(func() {
		TimeFieldFormat = "2006-01-02 15:04:05"
		AddHooks(func() meta.Meta {
			return meta.Int("socore", utils.RandInt(100))
		})

		newLog.Hook(func() meta.Meta {
			return meta.String("service", "demo")
		})

		newLog.Hook(func() meta.Meta {
			return meta.Int64("service-id", int64(utils.RandInt(1000)+1000))
		})

		newLog.Hook(func() meta.Meta {
			return meta.String("randomestr", utils.RandomString(1024))
		})
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newLog.Info("infomsg", meta.Int("uid", 12), meta.String("detail", "xxxxinfo...."))
	}

}

func TestStdLog(t *testing.T) {
	path := "testdata/overall.txt"
	file := openFile(path)
	defer file.Close()
	stdLog := newStdLogger(file)
	stdLog.Println(`{"time":"2020-08-07 11:09:58","level":"info","msg":"infomsg","uid":"12","detail":"xxxxinfo...."}`)
}

func BenchmarkStdLog(b *testing.B) {
	path := "testdata/overall.txt"
	file := openFile(path)
	defer file.Close()
	stdLog := newStdLogger(file)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf(`{"time":"%s","level":"%s","msg":"%s","uid":%d,"detail":"%s"}`,
			utils.TimeFormat(""), "infomsg", "detail", 12, "xxxxinfo....")
		stdLog.Println(s)
	}
}

func TestFileLog(t *testing.T) {
	path := "testdata/overall.txt"
	fl := newFileLog(path)
	defer fl.close()
	s := fmt.Sprintf(`{"time":"%s","level":"%s","msg":"%s","uid":%d,"detail":"%s"}`,
		utils.TimeFormat(""), "infomsg", "detail", 12, "xxxxinfo....")
	fl.write(s)
}

func BenchmarkFileLog(b *testing.B) {
	path := "testdata/overall.txt"
	fl := newFileLog(path)
	defer fl.close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("{\"time\":\"%s\",\"level\":\"%s\",\"msg\":\"%s\",\"uid\":%d,\"detail\":\"%s\"}\n",
			utils.TimeFormat(""), "infomsg", "detail", 12, "xxxxinfo....")
		fl.write(s)
	}
}

func TestZeroLog(t *testing.T) {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.MessageFieldName = "msg"
	path := "testdata/overall.txt"
	file := openFile(path)
	defer file.Close()
	logger := zerolog.New(file).With().Timestamp().Logger()
	logger.Info().Str("detail", "xxxxinfo....").Int("uid", 12).Msg("infomsg")
}

func BenchmarkZeroLog(b *testing.B) {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.MessageFieldName = "msg"
	path := "testdata/overall.txt"
	file := openFile(path)
	logger := zerolog.New(file).With().Timestamp().Logger()
	defer file.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().Str("detail", "xxxxinfo....").Int("uid", 12).Msg("infomsg")
	}
}

func newStdLogger(fi *os.File) *log.Logger {
	return log.New(fi, "", 0)
}

func openFile(path string) *os.File {
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	return file
}

type fileLog struct {
	file *os.File
	ch   []chan string
	sync.WaitGroup
}

func newFileLog(path string) *fileLog {
	file := openFile(path)
	f := &fileLog{
		file: file,
		ch:   make([]chan string, 10),
	}
	for i := range f.ch {
		f.ch[i] = make(chan string, 1024)
	}

	for i := range f.ch {
		f.Add(1)
		go func(ch chan string) {
			for {
				select {
				case s, ok := <-ch:
					if !ok {
						f.Done()
						return
					}
					f.file.WriteString(s)
				}
			}

		}(f.ch[i])
	}
	return f
}

func (f *fileLog) write(s string) {
	idx := utils.RandInt(10)
	f.ch[idx] <- s
}

func (f *fileLog) close() {
	for i := range f.ch {
		close(f.ch[i])
	}
	f.Wait()
	f.file.Close()
}

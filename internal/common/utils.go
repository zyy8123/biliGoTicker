package common

import (
	formatter "github.com/DaRealFreak/colored-nested-formatter"
	"github.com/beevik/ntp"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var (
	log  *Logger
	once sync.Once
)

type Logger struct {
	*logrus.Logger
	prefix string
}

func GetLogger(prefix string) *Logger {
	once.Do(func() {
		log = NewLogger(prefix)
	})
	return log
}

// NewLogger 返回带前缀的 Logger 实例
func NewLogger(prefix string) *Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		DisableColors:            false,
		ForceColors:              false,
		DisableTimestamp:         false,
		UseUppercaseLevel:        true,
		UseTimePassedAsTimestamp: false,
		TimestampFormat:          time.StampMilli,
		PadAllLogEntries:         true,
	})
	log = &Logger{Logger: l, prefix: prefix}
	return log
}

// customFormatter 实现 logrus.Formatter
type customFormatter struct {
	prefix string
}

func GetAccurateTime() time.Time {
	var ntpServers = []string{
		"ntp.aliyun.com",
		"cn.pool.ntp.org",
		"time.google.com",
		"time.windows.com",
		"pool.ntp.org",
	}
	for _, server := range ntpServers {
		resp, err := ntp.Query(server)
		if err != nil {
			log.Warningf("ntp %s 无法使用", server)
			continue
		}
		accurate := time.Now().Add(resp.ClockOffset)
		log.Infof("使用ntp %s,时间偏差 %s", server, resp.ClockOffset.String())
		return accurate
	}
	// 所有 NTP 失败，降级使用本地时间
	log.Errorf("所有 NTP 服务器都无法访问，使用本地时间。")
	return time.Now()
}
func SleepUntilAccurate(target time.Time) error {
	now := GetAccurateTime()
	if now.After(target) || now.Equal(target) {
		return nil
	}
	delta := target.Sub(now)
	cycleTimerChan := time.Tick(delta)
	select {
	case <-cycleTimerChan:
		return nil
	}
}

package infrastructure

import (
	"io"
	"os"
	"strconv"

	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	perDay, _ := strconv.Atoi(os.Getenv("LOG_ROTATOR_PER_DAY"))
	tw := &TimeWriter{
		Dir:           "log/",
		Compress:      true,
		ReserveDay:    perDay,
		LogFilePrefix: os.Getenv("APP_NAME"),
	}
	logger := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
	})
	logrus.SetOutput(io.MultiWriter(os.Stderr, tw))
	logrus.SetReportCaller(true)

	logrus.AddHook(&slackrus.SlackrusHook{
		HookURL:        os.Getenv("SLACK_WEBHOOK_URL"),
		AcceptedLevels: slackrus.LevelThreshold(logrus.ErrorLevel),
		Channel:        os.Getenv("SLACK_CHANNEL_NAME"),
		IconEmoji:      ":ghost:",
		Username:       "foobot",
	})
	return logger
}

package infrastructure

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Renos-id/go-starter-template/infrastructure/zapslack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLogger() *zap.Logger {
	perDay, _ := strconv.Atoi(os.Getenv("LOG_ROTATOR_PER_DAY"))
	w := zapcore.AddSync(&TimeWriter{
		Dir:           "log/",
		Compress:      true,
		ReserveDay:    perDay,
		LogFilePrefix: os.Getenv("APP_NAME"),
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		zap.InfoLevel,
	)
	var logger *zap.Logger
	//Check if Slack Webhook Url is set
	if os.Getenv("SLACK_WEBHOOK_URL") != "" {
		logger = zap.New(core, zap.AddCaller(), zap.Fields(), zap.Hooks(zapslack.NewSlackHook(os.Getenv("SLACK_WEBHOOK_URL"), zap.ErrorLevel, os.Getenv("APP_NAME")).GetHook()))
	} else {
		fmt.Println("No Slack")
		logger = zap.New(core)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	return logger
}

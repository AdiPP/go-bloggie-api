package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Renos-id/go-starter-template/infrastructure/httplog"
	"github.com/Renos-id/go-starter-template/infrastructure/slack"

	"github.com/rs/zerolog"
	"github.com/tidwall/gjson"
)

var errSkipEvent = errors.New("skip")

// CreateLogger creates a logger that logs everything to stderr.
// If the SENTRY_DSN environment variable is provided, it also sends events reported on error level to Sentry.
// Every event logged on error level will get unmarshaled and transformed into a Raven packet
// to be sent to Sentry.
func CreateLogger(lvl string) (zerolog.Logger, io.Closer, error) {
	tw := &httplog.TimeWriter{
		Dir:           "log/",
		Compress:      true,
		ReserveDay:    1,
		LogFilePrefix: os.Getenv("APP_NAME"),
	}

	pr, pw := io.Pipe()

	go func() {
		dec := json.NewDecoder(pr)

		for {
			var e logEvent
			err := dec.Decode(&e)

			if err == errSkipEvent {
				continue
			}

			if err != nil {
				fmt.Fprintf(os.Stderr, "unmarshaling log failed with error %v\n", err)
				continue
			}

		}
	}()

	// setup a global function that transforms any error passed to
	// zerolog to an error with stack strace.
	// zerolog.ErrorMarshalFunc = func(err error) interface{} {}

	return newLogger(lvl, io.MultiWriter(os.Stderr, pw, tw)), pw, nil
}

type logEvent struct{}

// unmarshal only if the level is error.
func (l *logEvent) UnmarshalJSON(data []byte) error {
	res := gjson.Get(string(data), "level")
	if !res.Exists() || res.String() != "error" {
		return errSkipEvent
	}
	sendToSlack(data)
	type event logEvent
	return json.Unmarshal(data, (*event)(l))
}

func sendToSlack(json_string []byte) {
	var mapped map[string]interface{}
	json.Unmarshal([]byte(json_string), &mapped)
	attachment1 := slack.Attachment{}
	for key, el := range mapped {
		attachment1.AddField(slack.Field{
			Title: key,
			Value: fmt.Sprint(el),
		})
	}
	payload := slack.Payload{
		Text:        fmt.Sprint(mapped["message"]),
		Username:    "robot",
		Channel:     os.Getenv("SLACK_CHANNEL_NAME"),
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.SendSlackMessage(os.Getenv("SLACK_WEBHOOK_URL"), payload)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

// newLogger returns a configured logger.
func newLogger(level string, w io.Writer) zerolog.Logger {
	logger := zerolog.New(w).With().Str("service", os.Getenv("APP_NAME")).Timestamp().Logger()

	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Fatal(err)
	}

	logger = logger.Level(lvl)

	// replace standard logger with zerolog
	log.SetFlags(0)
	log.SetOutput(logger)

	return logger
}

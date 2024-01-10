package logger

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlogInfoLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	slogHandler := slog.NewJSONHandler(&buffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLog := slog.New(slogHandler)
	log, _ := NewSlogLogger(slogLog)
	log.WithFields(Fields{
		"foo": "bar",
	}).Info("direct")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "INFO", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
}

func TestSlogInfofLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	slogHandler := slog.NewJSONHandler(&buffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLog := slog.New(slogHandler)
	log, _ := NewSlogLogger(slogLog)
	log.WithFields(Fields{
		"ping": "pong",
	}).Infof("received %s balls", "ping pong")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "received ping pong balls", fields["msg"])
	assert.Equal(t, "INFO", fields["level"])
	assert.Equal(t, "pong", fields["ping"])
}

func TestSlogWarnLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	slogHandler := slog.NewJSONHandler(&buffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLog := slog.New(slogHandler)
	log, _ := NewSlogLogger(slogLog)
	log.WithFields(Fields{
		"foo": "bar",
		"log": "slog",
	}).Warn("direct")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "direct", fields["msg"])
	assert.Equal(t, "WARN", fields["level"])
	assert.Equal(t, "bar", fields["foo"])
	assert.Equal(t, "slog", fields["log"])
}

func TestSlogWarnfLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	slogHandler := slog.NewJSONHandler(&buffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLog := slog.New(slogHandler)
	log, _ := NewSlogLogger(slogLog)
	log.WithFields(Fields{
		"ping": "pong",
		"log":  "slog",
	}).Warnf("received %s balls", "table tennis")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "received table tennis balls", fields["msg"])
	assert.Equal(t, "WARN", fields["level"])
	assert.Equal(t, "pong", fields["ping"])
	assert.Equal(t, "slog", fields["log"])
}

func TestSlogPanicLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	slogHandler := slog.NewJSONHandler(&buffer, &slog.HandlerOptions{
		Level: slog.LevelError,
	})
	slogLog := slog.New(slogHandler)
	log, _ := NewSlogLogger(slogLog)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
		err := json.Unmarshal(buffer.Bytes(), &fields)
		assert.Nil(t, err)
		assert.Equal(t, "db not found", fields["msg"])
		assert.Equal(t, "ERROR", fields["level"])
		assert.Equal(t, "dataDB", fields["db"])
		assert.Equal(t, "slog", fields["log"])
	}()
	log.WithFields(Fields{
		"db":  "dataDB",
		"log": "slog",
	}).Panic("db not found")
}

func TestSlogErrorLogger(t *testing.T) {
	var fields Fields
	var buffer bytes.Buffer
	slogHandler := slog.NewJSONHandler(&buffer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slogLog := slog.New(slogHandler)
	log, _ := NewSlogLogger(slogLog)
	log.WithFields(Fields{
		"acctNumber": 7899,
		"log":        "slog",
	}).Errorf("Error creating account %s", "testAccount")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)
	assert.Equal(t, "Error creating account testAccount", fields["msg"])
	assert.Equal(t, "ERROR", fields["level"])
	assert.Equal(t, float64(7899), fields["acctNumber"])
	assert.Equal(t, "slog", fields["log"])
}

// set logger to info and see that it doesn't print debug statements
func TestSlogNoOutputLogger(t *testing.T) {
	var buffer bytes.Buffer
	slogHandler := slog.NewJSONHandler(&buffer, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	slogLog := slog.New(slogHandler)
	log, _ := NewSlogLogger(slogLog)
	log.WithFields(Fields{
		"foo": "bar",
	}).Debugf("direct")

	assert.Equal(t, "", string(buffer.Bytes()))
}

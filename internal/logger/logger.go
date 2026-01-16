package logger

import (
	"time"
	"encoding/json"
	"os"
)

type Event struct {
	Timestamp time.Time `json:"timestamp"`

  SrcIP   string `json:"src_ip"`
  SrcPort int    `json:"src_port"`

  DstPort  int    `json:"dst_port"`
  Protocol string `json:"protocol"`

  RawPayload string `json:"raw_payload"`
}

type Logger interface {
    Log(event Event) error
}


type StdoutLogger struct{}

func NewStdoutLogger() *StdoutLogger {
    return &StdoutLogger{}
}

func (l *StdoutLogger) Log(event Event) error {
    event.Timestamp = time.Now()

    enc := json.NewEncoder(os.Stdout)
    return enc.Encode(event)
}


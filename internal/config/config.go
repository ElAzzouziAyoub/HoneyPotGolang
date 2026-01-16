package config

type Listener struct {
    Port     int    `json:"port"`
    Protocol string `json:"protocol"`
}

type Limits struct {
    MaxPayloadBytes int `json:"max_payload_bytes"`
    TimeoutSeconds  int `json:"timeout_seconds"`
}

type Config struct {
    Listeners []Listener `json:"listeners"`
    Limits    Limits     `json:"limits"`
}


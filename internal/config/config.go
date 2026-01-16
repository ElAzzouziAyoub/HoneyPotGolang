package config


import (
    "encoding/json"
    "fmt"
    "os"
)

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


func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }

    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("invalid config format: %w", err)
    }

    return &cfg, nil
}



func (c *Config) Validate() error {
    if len(c.Listeners) == 0 {
        return fmt.Errorf("no listeners defined")
    }

    ports := make(map[int]bool)

    for _, l := range c.Listeners {
        if l.Port <= 0 || l.Port > 65535 {
            return fmt.Errorf("invalid port: %d", l.Port)
        }

        if l.Protocol == "" {
            return fmt.Errorf("listener on port %d has empty protocol", l.Port)
        }

        if ports[l.Port] {
            return fmt.Errorf("duplicate port: %d", l.Port)
        }
        ports[l.Port] = true
    }

    if c.Limits.MaxPayloadBytes <= 0 {
        return fmt.Errorf("max_payload_bytes must be > 0")
    }

    if c.Limits.TimeoutSeconds <= 0 {
        return fmt.Errorf("timeout_seconds must be > 0")
    }

    return nil
}


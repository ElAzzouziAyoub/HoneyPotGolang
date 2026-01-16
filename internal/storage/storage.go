package storage

import (
    "encoding/json"
    "os"
    "honeypot/internal/logger"
)

type Storage interface {
    Save(event logger.Event) error
}
type FileStorage struct {
    Path string
}
func NewFileStorage(path string) *FileStorage {
    return &FileStorage{Path: path}
}
func (fs *FileStorage) Save(event logger.Event) error {
    file, err := os.OpenFile(fs.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    enc := json.NewEncoder(file)
    return enc.Encode(event)
}


package server

import (
    "bufio"
    "net"
    "log"

    "honeypot/internal/logger"
    "honeypot/internal/storage"
)

type Server struct {
    Addr    string
    Logger  logger.Logger
    Storage storage.Storage
}


func New(addr string, l logger.Logger, s storage.Storage) *Server {
    return &Server{
        Addr:    addr,
        Logger: l,
        Storage: s,
    }
}

func (s *Server) Start() {
    ln, err := net.Listen("tcp", s.Addr)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("listening on", s.Addr)

    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go s.handleConn(conn)
    }
}

func (s *Server) handleConn(conn net.Conn) {
    defer conn.Close()

    remoteAddr := conn.RemoteAddr().String()

    reader := bufio.NewReader(conn)
    data, err := reader.ReadString('\n')
    if err != nil {
        return
    }
		event := logger.Event{
        SrcIP:      remoteAddr,
        Protocol:   "tcp",
        RawPayload: data,
    }

    _ = s.Logger.Log(event)
    _ = s.Storage.Save(event)
	}


package main

import (
	"time"
)

const (
	tickrate = time.Second / 60
)

var server Server

type Server struct {
	EntityMap    map[uint16]EntityI
	NextEntityID uint32
	Tick         uint32
}

func StartServer() {
	server = Server{
		EntityMap:    make(map[uint16]EntityI),
		NextEntityID: 0,
	}

	ticker := time.NewTicker(tickrate)

	for {
		<-ticker.C

		for _, v := range server.EntityMap {
			next := v
			for next != nil {
				next.Step()
				next = next.GetNext()
			}
		}

		server.Tick++
	}
}

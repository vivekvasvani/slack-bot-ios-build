package main

import (
	"github.com/vivekvasvani/slack-bot-ios-build/server"
)

func main() {
	wait := make(chan struct{})
	server.NewServer()
	<-wait
}

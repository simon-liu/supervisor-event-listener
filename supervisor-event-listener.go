package main

import (
	"github.com/simon-liu/supervisor-event-listener/listener"
)

func main() {
	for {
		listener.Start()
	}
}

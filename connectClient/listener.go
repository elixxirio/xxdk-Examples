package main

import (
	jww "github.com/spf13/jwalterweatherman"

	"gitlab.com/elixxir/client/e2e/receive"
)

// listener implements the receive.Listener interface
type listener struct {
	name string
}

// Hear will be called whenever a message matching
// the RegisterListener call is received
// Message handling logic goes here
func (l listener) Hear(item receive.Message) {
	jww.INFO.Printf("Message received: %v", item)
}

// Name is used for debugging purposes
func (l listener) Name() string {
	return l.name
}

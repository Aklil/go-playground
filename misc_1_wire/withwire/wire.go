//go:build wireinject
// +build wireinject    

// the injector's purpose is to provide information about which providers to use to construct an Event and
// so we will exclude it from our final binary with a build constraint at the top of the file:
// using the above two lines 

package withwire

import "github.com/google/wire"

func InitializeEvent() Event {
  wire.Build(NewEvent, NewMessage, NewGreeter)
  return Event {}
}
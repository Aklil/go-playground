package main

import(
	withoutwire "aklil.com/wire_exercise/withoutwire"
	withwire "aklil.com/wire_exercise/withwire"
)

func initWithoutWire(){
	message := withoutwire.NewMessage()
    greeter := withoutwire.NewGreeter(message)
    event := withoutwire.NewEvent(greeter)

    event.Start()
}

func initWithWire() {
	e := withwire.InitializeEvent()
	e.Start()
}

func main() {

	initWithWire()

}
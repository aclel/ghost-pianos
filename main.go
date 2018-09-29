package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gomidi/connect"
	"github.com/gomidi/mid"
	driver "github.com/gomidi/rtmididrv"
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// This example expects the first input and output port to be connected
// somehow (are either virtual MIDI through ports or physically connected).
// We write to the out port and listen to the in port.
func main() {
	drv, err := driver.New()
	must(err)

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	must(err)

	outs, err := drv.Outs()
	must(err)

	if len(os.Args) == 2 && os.Args[1] == "list" {
		printInPorts(ins)
		printOutPorts(outs)
		return
	}

	in, out := ins[0], outs[0]

	must(in.Open())
	must(out.Open())

	wr := mid.WriteTo(out)

	// listen for MIDI
	go mid.NewReader().ReadFrom(in)

	{ // write MIDI to out that passes it to in on which we listen.
		wr.SetChannel(1)

		err := wr.NoteOn(60, 100)
		if err != nil {
			panic(err)
		}

		note := uint8(60)
		time.Sleep(time.Nanosecond)
		wr.NoteOff(note)
		time.Sleep(time.Nanosecond)
		wr.NoteOn(note, 100)
		time.Sleep(time.Second * 1)
		wr.NoteOff(note)
		time.Sleep(time.Nanosecond)

		if (note == 60) {
			fmt.Println("Reacting to note 60 - playing note 75")
			wr.NoteOn(75, 100)
			time.Sleep(time.Second * 1)
			wr.NoteOff(75)
			time.Sleep(time.Nanosecond)
		}
		

		// wr.NoteOn(70, 100)
		// time.Sleep(time.Nanosecond)
		// wr.NoteOff(70)
		// time.Sleep(time.Second * 1)
		// wr.NoteOn(60, 100)
		// time.Sleep(time.Nanosecond)
		// wr.NoteOff(60)
		// time.Sleep(time.Nanosecond)
	}
}

func printPort(port connect.Port) {
	fmt.Printf("[%v] %s\n", port.Number(), port.String())
}

func printInPorts(ports []connect.In) {
	fmt.Printf("MIDI IN Ports\n")
	for _, port := range ports {
		printPort(port)
	}
	fmt.Printf("\n\n")
}

func printOutPorts(ports []connect.Out) {
	fmt.Printf("MIDI OUT Ports\n")
	for _, port := range ports {
		printPort(port)
	}
	fmt.Printf("\n\n")
}
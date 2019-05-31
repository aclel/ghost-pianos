package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aclel/ghost-pianos/bjorklund"
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

	// ./ghost-pianos list
	if len(os.Args) == 2 && os.Args[1] == "list" {
		printInPorts(ins)
		printOutPorts(outs)
		return
	}

	in, out := ins[0], outs[0]

	must(in.Open())
	must(out.Open())

	wr := mid.WriteTo(out)

	// rd.Msg.Each = func() to hook into each message
	rd := mid.NewReader()
	rd.Msg.Channel.NoteOn = func(p *mid.Position, channel, key, vel uint8) {
		if key == 70 {
			minNote := 70
			maxNote := 90
			rhythm := bjorklund.Bjorklund(30, 18)

			randSource := rand.NewSource(time.Now().UnixNano())
			rand.New(randSource)
			nextKey := uint8(rand.Intn(maxNote-minNote) + minNote)
			go playRhythm(wr, nextKey, rhythm)
		}
	}

	// listen for MIDI
	go rd.ReadFrom(in)
	{
		for {

		}
	}
}

func playRhythm(wr *mid.Writer, note uint8, rhythm []int) {
	for i := 0; i < len(rhythm); i++ {
		pulse := rhythm[i]

		if pulse == 1 {
			wr.NoteOn(note, 100)
			time.Sleep(time.Second / 2)
			wr.NoteOff(note)
		} else {
			time.Sleep(time.Second / 2)
		}
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

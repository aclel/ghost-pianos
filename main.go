package main

import (
	"flag"
	"fmt"

	"github.com/gomidi/connect"
	"github.com/gomidi/mid"
	driver "github.com/gomidi/rtmididrv"
)

func main() {
	drv, err := driver.New()
	if err != nil {
		panic(err.Error())
	}
	defer drv.Close()

	ins, err := drv.Ins()
	if err != nil {
		panic(err.Error())
	}

	outs, err := drv.Outs()
	if err != nil {
		panic(err.Error())
	}

	if len(ins) < 1 {
		panic("There are no MIDI input ports available")
	}

	if len(outs) < 1 {
		panic("There are no MIDI output ports available")
	}

	inputPortPtr := flag.Int("in", 0, "MIDI input port")
	outputPortPtr := flag.Int("out", 0, "MIDI output port")
	bpmPtr := flag.Int("bpm", 120, "Beats per minute of the generated rhythms")
	velocityMultiplierPtr := flag.Float64("vel", 1, "Multiplier for output velocity")
	listPtr := flag.Bool("list", false, "List all of the available input and output MIDI ports")
	flag.Parse()

	if *listPtr {
		printInPorts(ins)
		printOutPorts(outs)
		return
	}

	selectedIn := *inputPortPtr
	if selectedIn > len(ins)-1 || selectedIn < 0 {
		panic("Selected input port: %v is out of range")
	}

	selectedOut := *outputPortPtr
	if selectedOut > len(outs)-1 || selectedOut < 0 {
		panic("Selected input port: %v is out of range")
	}

	in, out := ins[selectedIn], outs[selectedOut]

	err = in.Open()
	if err != nil {
		panic(err.Error())
	}

	err = out.Open()
	if err != nil {
		panic(err.Error())
	}

	writer := mid.WriteTo(out)
	noteGenerator := NoteGenerator{Writer: writer, BPM: *bpmPtr, VelocityMultiplier: *velocityMultiplierPtr}

	reader := mid.NewReader()
	reader.Msg.Channel.NoteOn = noteGenerator.RespondToKey

	// Indefinitely listen and respond MIDI inputs using the note generator
	go reader.ReadFrom(in)
	{
		for {

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

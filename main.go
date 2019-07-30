package main

import (
	"flag"
	"fmt"
	"io/ioutil"
  	"path/filepath"

	"github.com/aclel/ghost-pianos/bjorklund"
	"github.com/gomidi/connect"
	"github.com/gomidi/mid"
	"gopkg.in/yaml.v2"
	driver "github.com/gomidi/rtmididrv"
)

type Config struct {
	Rhythms map[uint8][]int `yaml:"rhythms"`
}

func (c* Config) buildRhythms() map[uint8][]int {
	rhythms := make(map[uint8][]int)
	for key, value := range c.Rhythms {
		rhythms[key] = bjorklund.Bjorklund(value[0], value[1])
	}
	return rhythms
}

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
	rhythmsPathPtr := flag.String("rhythms", "rhythms.yaml", "Relative path to a rhythms yaml file")
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

	filename, _ := filepath.Abs(*rhythmsPathPtr)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	rhythms := config.buildRhythms()

	writer := mid.WriteTo(out)
	noteGenerator := NoteGenerator{
		Rhythms: rhythms,
		Writer: writer,
		BPM: *bpmPtr,
		VelocityMultiplier: *velocityMultiplierPtr,
	}

	reader := mid.NewReader()
	reader.Msg.Channel.NoteOn = noteGenerator.RespondToKeyPlusSevenRando

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

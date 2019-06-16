package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gomidi/mid"
)

type NoteGeneratorInterface interface {
	RespondToKey(p *mid.Position, channel, key, velocity uint8)
	RespondToKeyPlusSevenRando(p *mid.Position, channel, key, velocity uint8)
	PlayRhythm(note uint8, rhythm []int, velocity uint8)
}

type NoteGenerator struct {
	Writer *mid.Writer
	BPM    int
}

func (noteGenerator NoteGenerator) RespondToKey(p *mid.Position, channel, key, velocity uint8) {
	rhythm := Rhythms[key]
	sleepDuration := convertBPMToMilliSeconds(noteGenerator.BPM)
	fmt.Printf("[%v] %v\n", key, sleepDuration)
	go noteGenerator.PlayRhythm(key, rhythm, velocity)
}

func (noteGenerator NoteGenerator) RespondToKeyPlusSevenRando(p *mid.Position, channel, key, velocity uint8) {
	rhythm := Rhythms[key]

	randSource := rand.NewSource(time.Now().UnixNano())
	rand.New(randSource)

	operators := []int{1, -1}
	operator := operators[rand.Intn(len(operators))]

	newKey := int(key) + (7 * operator)
	if newKey > 88 {
		newKey = 88
	} else if newKey < 1 {
		newKey = 1
	}

	fmt.Printf("[%v] %v\n", newKey, key)
	go noteGenerator.PlayRhythm(uint8(newKey), rhythm, velocity)
}

func (noteGenerator NoteGenerator) PlayDefinedSequence(p *mid.Position, channel, key, velocity uint8) {
	fmt.Printf("[%v]\n", key)
	go noteGenerator.PlayNoteSequence(key, 50)
}

func (noteGenerator NoteGenerator) PlayRhythm(note uint8, rhythm []int, velocity uint8) {
	sleepDuration := convertBPMToMilliSeconds(noteGenerator.BPM)
	for i := 0; i < len(rhythm); i++ {
		pulse := rhythm[i]

		if pulse == 1 {
			noteGenerator.Writer.NoteOn(note, velocity)
			time.Sleep(sleepDuration)
			noteGenerator.Writer.NoteOff(note)
		} else {
			time.Sleep(sleepDuration)
		}
	}
}

func (noteGenerator NoteGenerator) PlayNoteSequence(key uint8, velocity uint8) {
	sleepDuration := convertBPMToMilliSeconds(noteGenerator.BPM)

	keyInt := int(key)
	if keyInt > len(Sequences)-1 {
		time.Sleep(sleepDuration)
	}
	notes := Sequences[keyInt]

	for i := 0; i < len(notes); i++ {
		note := notes[i]
		if i%2 == 0 || i%3 == 0 {
			velocity = 100
		} else {
			velocity = 70
		}
		if note > 0 {
			noteGenerator.Writer.NoteOn(note, velocity)
			time.Sleep(sleepDuration)
			noteGenerator.Writer.NoteOff(note)
		} else {
			time.Sleep(sleepDuration)
		}
	}
}

func convertBPMToMilliSeconds(bpm int) time.Duration {
	milliSeconds := 60000 / bpm
	return time.Duration(milliSeconds) * time.Millisecond
}

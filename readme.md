# Ghost Pianos

## Prerequisites

* vmpk for midi input

## Getting started

1. Run two instances of vmpk with `open -n /Applications/vmpk.app` - ensure that MIDI Input and output are set to Core Audio
2. Run ghost-pianos with `-in` set to the input port of one instance VMPK (the Living Piano) and `-out` set to the output port of the other instance (the Ghost Piano).
3. Run a DAW and hook up a track to the output from the Ghost Piano. You should be able to hear the ghosts.

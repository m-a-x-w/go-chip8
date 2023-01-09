package main

import (
	"math/rand"
	"time"
)

var (
	MemStartAddress = uint16(0x20)
	MemFontsetStart = uint16(0x50)
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	emu := Chip8{
		registers:  [16]uint8{},
		memory:     [4096]uint8{},
		index:      0,
		pc:         MemStartAddress,
		stack:      [16]uint16{},
		sp:         0,
		delayTimer: 0,
		soundTimer: 0,
		keypad:     [16]uint8{},
		video:      [2048]uint32{},
		opcode:     0,
	}

	emu.loadRom("test_opcode.ch8")

	for i := uint16(0); i < MemFontsetStart; i++ {
		emu.memory[MemFontsetStart+i] = fontset[i]
	}
}

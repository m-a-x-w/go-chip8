package main

import (
	"math/rand"
	"os"
)

type Chip8 struct {
	registers  [16]uint8
	memory     [4096]uint8
	index      uint16
	pc         uint16
	stack      [16]uint16
	sp         uint8
	delayTimer uint8
	soundTimer uint8
	keypad     [16]uint8
	video      [64 * 32]uint32
	opcode     uint16
}

func (c *Chip8) loadRom(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	for i := uint16(0); i < (uint16)(size); i++ {
		c.memory[MemStartAddress+i] = buffer[i]
	}

}
func (c *Chip8) randByte() byte {
	return byte(rand.Intn(256))
}

func (c *Chip8) Op00e0() { // CLS: clear display
	c.video = [64 * 32]uint32{}
}

func (c *Chip8) Op00ee() { // RET: return from a subroutine
	c.sp--
	c.pc = c.stack[c.sp]
}

func (c *Chip8) Op1nnn() { // JP addr: Jump to location nnn, set pc to nnn
	c.pc = c.opcode & 0x0FFF
}

func (c *Chip8) Op2nnn() { // CALL addr: Call subroutine at nnn
	address := c.opcode & 0x0FFF
	c.stack[c.sp] = c.pc
	c.sp++
	c.pc = address
}

func (c *Chip8) Op3xkk() { // SE Vx, byte: Skip next if Vx == kk
	Vx := (c.opcode & 0x0F00) >> 8
	bytee := uint8((c.opcode & 0x0F00) >> 8)

	if c.registers[Vx] == bytee {
		c.pc += 2
	}
}

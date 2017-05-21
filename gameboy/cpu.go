package gameboy

import (
	"strconv"
	"fmt"
)

/*
	gameboy central processing unit
 */
type CPU struct {

	Gameboy		*Gameboy;

	register_af  	uint16;
	register_bc  	uint16;
	register_de  	uint16;
	register_hl 	uint16;
	register_sp  	uint16;
	register_pc  	uint16;

	opcode	  	[256]func();
	initOpCodeArray func();

}

/*//Constructor
func (gameboy *Gameboy) NewCPU() CPU {
	cpu := CPU{
		Gameboy: gameboy,
		register_af: uint16(0x00),
		register_bc: uint16(0x00),
		register_de: uint16(0x00),
		register_hl: uint16(0x00),
		register_sp: uint16(0x00),
		register_pc: uint16(0x00),
	};
	cpu.initOpcodeArray();
	return cpu;
}*/

//HELPER FUNCTION FOR BINARY NUMBERS
func BtoInt64(s string) int64 {
	if i, err := strconv.ParseInt(s, 2, 64); err != nil {
		fmt.Println(err);
		return 0x00;
	} else {
		return i;
	}
}
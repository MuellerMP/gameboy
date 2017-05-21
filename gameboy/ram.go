package gameboy


type RAM struct {
	Gameboy		*Gameboy;

}

/*//Constructor
func (gameboy *Gameboy) NewRAM() RAM {
	ram := RAM{
		Gameboy: gameboy,
	};
	return ram;
}*/

// read value from 16bit Adress
func (ram *RAM) ReadMem(address uint16) uint8 {
	return 0x00
}

func (ram *RAM) WriteMem(address uint16, value uint8) {
	//TODO: write to Memory
}

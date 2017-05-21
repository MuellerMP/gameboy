package gameboy


type Gameboy struct {
	CPU	CPU
	RAM 	RAM
}

//Constructor
func New() *Gameboy {
	gameboy := new(Gameboy)
	gameboy.CPU.Gameboy = gameboy;
	gameboy.RAM.Gameboy = gameboy;
	return gameboy
}
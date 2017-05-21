package gameboy

//	FLAGS

func (cpu *CPU) GetSignFlag() bool {
	return ( cpu.register_af & 0x0080 ) != 0;
}

func (cpu *CPU) GetZeroFlag() bool {
	return ( cpu.register_af & 0x0040 ) != 0;
}

func (cpu *CPU) GetParityFlag() bool {
	return ( cpu.register_af & 0x0004 ) != 0;
}

func (cpu *CPU) GetCarryFlag() bool {
	return ( cpu.register_af & 0x0001 ) != 0;
}


func (cpu *CPU) SetSignFlag() {
	cpu.register_af |= 0x0080;
}
func (cpu *CPU) ClearSignFlag() {
	cpu.register_af &= 0xFF7F;
}

func (cpu *CPU) SetZeroFlag() {
	cpu.register_af |= 0x0040;
}
func (cpu *CPU) ClearZeroFlag() {
	cpu.register_af &= 0xFFBF;
}

func (cpu *CPU) SetParityFlag() {
	cpu.register_af |= 0x0004;
}
func (cpu *CPU) ClearParityFlag() {
	cpu.register_af &= 0xFFFB;
}

func (cpu *CPU) SetCarryFlag() {
	cpu.register_af |= 0x0001;
}
func (cpu *CPU) ClearCarryFlag() {
	cpu.register_af &= 0xFFFE;
}

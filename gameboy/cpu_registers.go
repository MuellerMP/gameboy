package gameboy

//	REGISTERS
/*
	Read 16bit Regs
 */
func (cpu *CPU) ReadAF() uint16 {
	return cpu.register_af;
}

func (cpu *CPU) ReadBC() uint16 {
	return cpu.register_bc;
}

func (cpu *CPU) ReadDE() uint16 {
	return cpu.register_de;
}

func (cpu *CPU) ReadHL() uint16 {
	return cpu.register_hl;
}

func (cpu *CPU) ReadSP() uint16 {
	return cpu.register_sp;
}

func (cpu *CPU) ReadPC() uint16 {
	return cpu.register_pc;
}

/*
	Read 8bit Regs
 */
func (cpu *CPU) ReadA() uint8 {
	return uint8((cpu.register_af >> 8) & 0xFF);
}
func (cpu *CPU) ReadB() uint8 {
	return uint8((cpu.register_bc >> 8) & 0xFF);
}
func (cpu *CPU) ReadC() uint8 {
	return uint8(cpu.register_bc & 0xFF);
}
func (cpu *CPU) ReadD() uint8 {
	return uint8((cpu.register_de >> 8) & 0xFF);
}
func (cpu *CPU) ReadE() uint8 {
	return uint8(cpu.register_de & 0xFF);
}
func (cpu *CPU) ReadH() uint8 {
	return uint8((cpu.register_hl >> 8) & 0xFF);
}
func (cpu *CPU) ReadL() uint8 {
	return uint8(cpu.register_hl & 0xFF);
}


/*
	Write to 16bit Regs
 */
func (cpu *CPU) WriteAF(value uint16) {
	cpu.register_af = value;
}

func (cpu *CPU) WriteBC(value uint16)  {
	cpu.register_bc = value;
}

func (cpu *CPU) WriteDE(value uint16) {
	cpu.register_de = value;
}

func (cpu *CPU) WriteHL(value uint16) {
	cpu.register_hl = value;
}

func (cpu *CPU) WriteSP(value uint16) {
	cpu.register_sp = value;
}

func (cpu *CPU) WritePC(value uint16) {
	cpu.register_pc = value;
}

/*
	Write to 8bit Regs
 */
func (cpu *CPU) WriteA(value uint8) {
	cpu.register_af &= 0x00FF;
	cpu.register_af |= (uint16(value) << 8) & 0xFF00;
}
func (cpu *CPU) WriteB(value uint8) {
	cpu.register_bc &= 0x00FF;
	cpu.register_bc |= (uint16(value) << 8) & 0xFF00;
}
func (cpu *CPU) WriteC(value uint8) {
	cpu.register_bc &= 0xFF00;
	cpu.register_bc |= uint16(value) & 0x00FF;
}
func (cpu *CPU) WriteD(value uint8) {
	cpu.register_de &= 0x00FF;
	cpu.register_de |= (uint16(value) << 8) & 0xFF00;
}
func (cpu *CPU) WriteE(value uint8) {
	cpu.register_de &= 0xFF00;
	cpu.register_de |= uint16(value) & 0x00FF;
}
func (cpu *CPU) WriteH(value uint8) {
	cpu.register_hl &= 0x00FF;
	cpu.register_hl |= (uint16(value) << 8) & 0xFF00;
}
func (cpu *CPU) WriteL(value uint8) {
	cpu.register_hl &= 0xFF00;
	cpu.register_hl |= uint16(value) & 0x00FF;
}
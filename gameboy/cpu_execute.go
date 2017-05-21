package gameboy

//	OPCODE

func (cpu *CPU) initOpcodeArray() {

	cpu.opcode[BtoInt64("01000000")] = func(){cpu.moveBToB()};
	cpu.opcode[BtoInt64("01000001")] = func(){cpu.moveBToC()};
	cpu.opcode[BtoInt64("01000010")] = func(){cpu.moveBToD()};
	cpu.opcode[BtoInt64("01000011")] = func(){cpu.moveBToE()};
	cpu.opcode[BtoInt64("01000100")] = func(){cpu.moveBToH()};
	cpu.opcode[BtoInt64("01000101")] = func(){cpu.moveBToL()};
	cpu.opcode[BtoInt64("01000110")] = func(){cpu.moveBToRAM()};
	cpu.opcode[BtoInt64("01000111")] = func(){cpu.moveBToA()};

	cpu.opcode[BtoInt64("01001000")] = func(){cpu.moveCToB()};
	cpu.opcode[BtoInt64("01001001")] = func(){cpu.moveCToC()};
	cpu.opcode[BtoInt64("01001010")] = func(){cpu.moveCToD()};
	cpu.opcode[BtoInt64("01001011")] = func(){cpu.moveCToE()};
	cpu.opcode[BtoInt64("01001100")] = func(){cpu.moveCToH()};
	cpu.opcode[BtoInt64("01001101")] = func(){cpu.moveCToL()};
	cpu.opcode[BtoInt64("01001110")] = func(){cpu.moveCToRAM()};
	cpu.opcode[BtoInt64("01001111")] = func(){cpu.moveCToA()};

	cpu.opcode[BtoInt64("01010000")] = func(){cpu.moveDToB()};
	cpu.opcode[BtoInt64("01010001")] = func(){cpu.moveDToC()};
	cpu.opcode[BtoInt64("01010010")] = func(){cpu.moveDToD()};
	cpu.opcode[BtoInt64("01010011")] = func(){cpu.moveDToE()};
	cpu.opcode[BtoInt64("01010100")] = func(){cpu.moveDToH()};
	cpu.opcode[BtoInt64("01010101")] = func(){cpu.moveDToL()};
	cpu.opcode[BtoInt64("01010110")] = func(){cpu.moveDToRAM()};
	cpu.opcode[BtoInt64("01010111")] = func(){cpu.moveDToA()};

	cpu.opcode[BtoInt64("01011000")] = func(){cpu.moveEToB()};
	cpu.opcode[BtoInt64("01011001")] = func(){cpu.moveEToC()};
	cpu.opcode[BtoInt64("01011010")] = func(){cpu.moveEToD()};
	cpu.opcode[BtoInt64("01011011")] = func(){cpu.moveEToE()};
	cpu.opcode[BtoInt64("01011100")] = func(){cpu.moveEToH()};
	cpu.opcode[BtoInt64("01011101")] = func(){cpu.moveEToL()};
	cpu.opcode[BtoInt64("01011110")] = func(){cpu.moveEToRAM()};
	cpu.opcode[BtoInt64("01011111")] = func(){cpu.moveEToA()};

	cpu.opcode[BtoInt64("01100000")] = func(){cpu.moveHToB()};
	cpu.opcode[BtoInt64("01100001")] = func(){cpu.moveHToC()};
	cpu.opcode[BtoInt64("01100010")] = func(){cpu.moveHToD()};
	cpu.opcode[BtoInt64("01100011")] = func(){cpu.moveHToE()};
	cpu.opcode[BtoInt64("01100100")] = func(){cpu.moveHToH()};
	cpu.opcode[BtoInt64("01100101")] = func(){cpu.moveHToL()};
	cpu.opcode[BtoInt64("01100110")] = func(){cpu.moveHToRAM()};
	cpu.opcode[BtoInt64("01100111")] = func(){cpu.moveHToA()};

	cpu.opcode[BtoInt64("01101000")] = func(){cpu.moveLToB()};
	cpu.opcode[BtoInt64("01101001")] = func(){cpu.moveLToC()};
	cpu.opcode[BtoInt64("01101010")] = func(){cpu.moveLToD()};
	cpu.opcode[BtoInt64("01101011")] = func(){cpu.moveLToE()};
	cpu.opcode[BtoInt64("01101100")] = func(){cpu.moveLToH()};
	cpu.opcode[BtoInt64("01101101")] = func(){cpu.moveLToL()};
	cpu.opcode[BtoInt64("01101110")] = func(){cpu.moveLToRAM()};
	cpu.opcode[BtoInt64("01101111")] = func(){cpu.moveLToA()};

	cpu.opcode[BtoInt64("01110000")] = func(){cpu.moveRAMToB()};
	cpu.opcode[BtoInt64("01110001")] = func(){cpu.moveRAMToC()};
	cpu.opcode[BtoInt64("01110010")] = func(){cpu.moveRAMToD()};
	cpu.opcode[BtoInt64("01110011")] = func(){cpu.moveRAMToE()};
	cpu.opcode[BtoInt64("01110100")] = func(){cpu.moveRAMToH()};
	cpu.opcode[BtoInt64("01110101")] = func(){cpu.moveRAMToL()};
	cpu.opcode[BtoInt64("01110110")] = func(){cpu.moveRAMToRAM()};
	cpu.opcode[BtoInt64("01110111")] = func(){cpu.moveRAMToA()};

	cpu.opcode[BtoInt64("01111000")] = func(){cpu.moveAToB()};
	cpu.opcode[BtoInt64("01111001")] = func(){cpu.moveAToC()};
	cpu.opcode[BtoInt64("01111010")] = func(){cpu.moveAToD()};
	cpu.opcode[BtoInt64("01111011")] = func(){cpu.moveAToE()};
	cpu.opcode[BtoInt64("01111100")] = func(){cpu.moveAToH()};
	cpu.opcode[BtoInt64("01111101")] = func(){cpu.moveAToL()};
	cpu.opcode[BtoInt64("01111110")] = func(){cpu.moveAToRAM()};
	cpu.opcode[BtoInt64("01111111")] = func(){cpu.moveAToA()};

}

func (cpu *CPU) ExecuteInstruction() {
	var next_opcode uint8 = 0;
	cpu.register_pc++;
	cpu.opcode[next_opcode]();
}


//CPU Operations: DATA TRANSFER GROUP
//REG A
func (cpu *CPU) moveAToA() {
	//NOOP
}
func (cpu *CPU) moveAToB() {
	cpu.WriteA(cpu.ReadA());
}
func (cpu *CPU) moveAToC() {
	cpu.WriteA(cpu.ReadC());
}
func (cpu *CPU) moveAToD() {
	cpu.WriteA(cpu.ReadC());
}
func (cpu *CPU) moveAToE() {
	cpu.WriteA(cpu.ReadE());
}
func (cpu *CPU) moveAToH() {
	cpu.WriteA(cpu.ReadH());
}
func (cpu *CPU) moveAToL() {
	cpu.WriteA(cpu.ReadL());
}
func (cpu *CPU) moveAToRAM() {
	cpu.Gameboy.RAM.WriteMem(cpu.ReadHL(), cpu.ReadA());
}

//REG B
func (cpu *CPU) moveBToA() {
	cpu.WriteB(cpu.ReadA());
}
func (cpu *CPU) moveBToB() {
	//NOOP
}
func (cpu *CPU) moveBToC() {
	cpu.WriteB(cpu.ReadC());
}
func (cpu *CPU) moveBToD() {
	cpu.WriteB(cpu.ReadD());
}
func (cpu *CPU) moveBToE() {
	cpu.WriteB(cpu.ReadE());
}
func (cpu *CPU) moveBToH() {
	cpu.WriteB(cpu.ReadH());
}
func (cpu *CPU) moveBToL() {
	cpu.WriteB(cpu.ReadL());
}
func (cpu *CPU) moveBToRAM() {
	cpu.Gameboy.RAM.WriteMem(cpu.ReadHL(), cpu.ReadB());
}

//REG C
func (cpu *CPU) moveCToA() {
	cpu.WriteC(cpu.ReadA());
}
func (cpu *CPU) moveCToB() {
	cpu.WriteC(cpu.ReadB());
}
func (cpu *CPU) moveCToC() {
	//NOOP
}
func (cpu *CPU) moveCToD() {
	cpu.WriteC(cpu.ReadD());
}
func (cpu *CPU) moveCToE() {
	cpu.WriteC(cpu.ReadE());
}
func (cpu *CPU) moveCToH() {
	cpu.WriteC(cpu.ReadH());
}
func (cpu *CPU) moveCToL() {
	cpu.WriteC(cpu.ReadL());
}
func (cpu *CPU) moveCToRAM() {
	cpu.Gameboy.RAM.WriteMem(cpu.ReadHL(), cpu.ReadC());
}

//REG D
func (cpu *CPU) moveDToA() {
	cpu.WriteD(cpu.ReadA());
}
func (cpu *CPU) moveDToB() {
	cpu.WriteD(cpu.ReadB());
}
func (cpu *CPU) moveDToC() {
	cpu.WriteD(cpu.ReadA());
}
func (cpu *CPU) moveDToD() {
	//NOOP
}
func (cpu *CPU) moveDToE() {
	cpu.WriteD(cpu.ReadE());
}
func (cpu *CPU) moveDToH() {
	cpu.WriteD(cpu.ReadH());
}
func (cpu *CPU) moveDToL() {
	cpu.WriteD(cpu.ReadL());
}
func (cpu *CPU) moveDToRAM() {
	cpu.Gameboy.RAM.WriteMem(cpu.ReadHL(), cpu.ReadD());
}

//REG E
func (cpu *CPU) moveEToA() {
	cpu.WriteE(cpu.ReadA());
}
func (cpu *CPU) moveEToB() {
	cpu.WriteE(cpu.ReadB());
}
func (cpu *CPU) moveEToC() {
	cpu.WriteE(cpu.ReadC());
}
func (cpu *CPU) moveEToD() {
	cpu.WriteE(cpu.ReadD());
}
func (cpu *CPU) moveEToE() {
	//NOOP
}
func (cpu *CPU) moveEToH() {
	cpu.WriteE(cpu.ReadH());
}
func (cpu *CPU) moveEToL() {
	cpu.WriteE(cpu.ReadL());
}
func (cpu *CPU) moveEToRAM() {
	cpu.Gameboy.RAM.WriteMem(cpu.ReadHL(), cpu.ReadE());
}

//REG H
func (cpu *CPU) moveHToA() {
	cpu.WriteH(cpu.ReadA());
}
func (cpu *CPU) moveHToB() {
	cpu.WriteH(cpu.ReadB());
}
func (cpu *CPU) moveHToC() {
	cpu.WriteH(cpu.ReadC());
}
func (cpu *CPU) moveHToD() {
	cpu.WriteH(cpu.ReadD());
}
func (cpu *CPU) moveHToE() {
	cpu.WriteH(cpu.ReadE());
}
func (cpu *CPU) moveHToH() {
	//NOOP
}
func (cpu *CPU) moveHToL() {
	cpu.WriteH(cpu.ReadL());
}
func (cpu *CPU) moveHToRAM() {
	cpu.Gameboy.RAM.WriteMem(cpu.ReadHL(), cpu.ReadH());
}

//REG L
func (cpu *CPU) moveLToA() {
	cpu.WriteL(cpu.ReadA());
}
func (cpu *CPU) moveLToB() {
	cpu.WriteL(cpu.ReadB());
}
func (cpu *CPU) moveLToC() {
	cpu.WriteL(cpu.ReadC());
}
func (cpu *CPU) moveLToD() {
	cpu.WriteL(cpu.ReadD());
}
func (cpu *CPU) moveLToE() {
	cpu.WriteL(cpu.ReadE());
}
func (cpu *CPU) moveLToH() {
	cpu.WriteL(cpu.ReadH());
}
func (cpu *CPU) moveLToL() {
	//NOOP
}
func (cpu *CPU) moveLToRAM() {
	cpu.Gameboy.RAM.WriteMem(cpu.ReadHL(), cpu.ReadL());
}


//RAM
func (cpu *CPU) moveRAMToA() {
	address := cpu.ReadHL();
	cpu.WriteA(cpu.Gameboy.RAM.ReadMem(address));
}
func (cpu *CPU) moveRAMToB() {
	address := cpu.ReadHL();
	cpu.WriteB(cpu.Gameboy.RAM.ReadMem(address));
}
func (cpu *CPU) moveRAMToC() {
	address := cpu.ReadHL();
	cpu.WriteC(cpu.Gameboy.RAM.ReadMem(address));
}
func (cpu *CPU) moveRAMToD() {
	address := cpu.ReadHL();
	cpu.WriteD(cpu.Gameboy.RAM.ReadMem(address));
}
func (cpu *CPU) moveRAMToE() {
	address := cpu.ReadHL();
	cpu.WriteE(cpu.Gameboy.RAM.ReadMem(address));
}
func (cpu *CPU) moveRAMToH() {
	address := cpu.ReadHL();
	cpu.WriteH(cpu.Gameboy.RAM.ReadMem(address));
}
func (cpu *CPU) moveRAMToL() {
	address := cpu.ReadHL();
	cpu.WriteL(cpu.Gameboy.RAM.ReadMem(address));
}
func (cpu *CPU) moveRAMToRAM() {
	//NOOP
}

//TODO: ARTITHMETIC GROUP
//TODO: LOGICAL GROUP

//CPU Operations: BRANCH GROUP
//Regular Jump
func (cpu *CPU) jump(){

}
//Conditional Jumps
func (cpu *CPU) jumpNotZero(){

}
func (cpu *CPU) jumpZero(){

}
func (cpu *CPU) jumpNotCarry(){

}
func (cpu *CPU) jumpCarry(){

}
func (cpu *CPU) jumpParityOdd(){

}
func (cpu *CPU) jumpParityEven(){

}
func (cpu *CPU) jumpPlus(){

}
func (cpu *CPU) jumpMinus(){

}
//Conditional Calls

func (cpu *CPU) callNotZero(){

}
func (cpu *CPU) callZero(){

}
func (cpu *CPU) callNotCarry(){

}
func (cpu *CPU) callCarry(){

}
func (cpu *CPU) callParityOdd(){

}
func (cpu *CPU) callParityEven(){

}
func (cpu *CPU) callPlus(){

}
func (cpu *CPU) callMinus(){

}
//RETURN
func (cpu *CPU) ret(){

}
//Conditional Returns

func (cpu *CPU) returnNotZero(){

}
func (cpu *CPU) returnZero(){

}
func (cpu *CPU) returnNotCarry(){

}
func (cpu *CPU) returnCarry(){

}
func (cpu *CPU) returnParityOdd(){

}
func (cpu *CPU) returnParityEven(){

}
func (cpu *CPU) returnPlus(){

}
func (cpu *CPU) returnMinus(){

}

//Restart/s
func (cpu *CPU) restart000(){

}
func (cpu *CPU) restart001(){

}
func (cpu *CPU) restart010(){

}
func (cpu *CPU) restart011(){

}
func (cpu *CPU) restart100(){

}
func (cpu *CPU) restart101(){

}
func (cpu *CPU) restart110(){

}
func (cpu *CPU) restart111(){

}


func (cpu *CPU) jumpHL() {

}
//TODO: STACK, I/0, MACHINE CONTROL GROUP

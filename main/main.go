package main

import (
	"github.com/LokiTheMango/gameboy/gameboy"
	"fmt"
)

func main () {
	var gameboy = gameboy.New();
	fmt.Println(gameboy);
	fmt.Println("EVERTHING GUD");
	gameboy.CPU.ExecuteInstruction();
}
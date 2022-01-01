package main

import (
	"fmt"
	"bufio"
	"os"
)
func game(_firstNumber,_secondNumber,_subtraction,_answer){
	reader:=bufio.NewReader(os.Stdin)
	//display a welcome/intstructions
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take them through the game
	fmt.Println("Multiply your number by", _firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply result by", _secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result with the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", _subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer
	fmt.Println("The answer is", _answer)
}
package main

import (
	"bufio"
	"fmt"
	"os"
)
const prompt = "and press ENTER when ready"

func main() {
	var firstNumber =2 
	var secondNumber = 5
	var subtraction = 7
	// var answer int

	reader:=bufio.NewReader(os.Stdin)
	//display a welcome/intstructions
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take them through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result with the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer

}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)
const prompt = "and don't type your number in , just press ENTER when ready"

func main() {

	//seed the generator
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8)+2
	var secondNumber = rand.Intn(8)+2
	var subtraction = rand.Intn(8)+2
	var answer int

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
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}

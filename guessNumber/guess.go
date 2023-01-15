package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again")
			return
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input. Please enter an integer value")
			return
		}
		fmt.Println("Your guess is ", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than secret number")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than secret number")
		} else {
			fmt.Println("Correct, you Legend")
			break
		}
	}
}

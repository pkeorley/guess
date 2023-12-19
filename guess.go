package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	success := false
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100.\nCan you guess it?")

	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {

		fmt.Print(fmt.Sprintf("You gave %d guesses left.\nMake a guess: ", 10-guesses))
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}

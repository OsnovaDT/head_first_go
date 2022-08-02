// Игра, в которой пользователь может попробовать себя в угадывании числа

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func handle_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	reader := bufio.NewReader(os.Stdin)

	random_number := rand.Intn(101)
	max_guess_amount := 10
	is_guessed := false

	for guess_amount := 0; guess_amount < max_guess_amount; guess_amount++ {
		fmt.Println("\nYou have", max_guess_amount-guess_amount, "guesses")
		fmt.Print("Please guess a number from 1 to 100: ")

		user_guess, err := reader.ReadString('\n')
		handle_error(err)

		user_number, err := strconv.Atoi(strings.TrimSpace(user_guess))
		handle_error(err)

		if user_number < random_number {
			fmt.Println("Oops. Your guess was LOW")
		} else if user_number > random_number {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			fmt.Println("Yes, it's", random_number)

			is_guessed = true

			break
		}
	}

	if !is_guessed {
		fmt.Println("Sorry. You didn't guess my number", random_number)
	}
}

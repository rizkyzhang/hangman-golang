package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	figures := [7]string{
		`
  +---+
  0   |
 /|\  |
 / \  |
    =====
  `,
		`
  +---+
  0   |
 /|\  |
 /    |
    =====
  `,
		`
  +---+
  0   |
 /|\  |
      |
    =====
  `,
		`
  +---+
  0   |
 /|   |
      |
    =====
  `,
		`
  +---+
  0   |
  |   |
      |
    =====
  `,
		`
  +---+
  0   |
      |
      |
    =====
  `,
		`
  +---+
      |
      |
      |
    =====
  `,
	}

	life := 6
	randWord := strings.ToUpper(gofakeit.Word())
	correctLetters := strings.Split(strings.Repeat("_", len(randWord)), "")
	var wrongGuesses []string
	var guesses string

	for {
		// Figure and stats
		fmt.Println(figures[life])
		fmt.Println(strings.Join(correctLetters, ""))
		fmt.Printf("Wrong guesses: %s\n", wrongGuesses)

		// Lose
		if life == 0 {
			fmt.Println("You lose!")
			fmt.Printf("Random word: %s\n", randWord)
			fmt.Printf("Guesses: %v", guesses)
			break
		}

		// Win
		if strings.Join(correctLetters, "") == randWord {
			fmt.Println("Congrats! You win!")
			fmt.Printf("Random word: %s\n", randWord)
			fmt.Printf("Guesses: %v", guesses)
			break
		}

		// Get input from user
		fmt.Print("Please enter a letter: ")
		reader := bufio.NewReader(os.Stdin)
		_guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess := strings.ToUpper(strings.TrimSpace(_guess))

		// Validate input
		isLetter := unicode.IsLetter([]rune(guess)[0])
		if !isLetter || len(guess) > 1 {
			fmt.Println("Only letter is allowed")
			continue
		}

		// Check if letter has been guessed before
		if strings.Contains(guesses, guess) {
			fmt.Println("Letter has been guessed before")
			continue
		}
		guesses += guess

		// Check if letter is correct or wrong
		if strings.Contains(randWord, guess) {
			fmt.Println("Correct guess")
			correctLetters[strings.Index(randWord, guess)] = guess
		} else {
			fmt.Println("Wrong guess")
			wrongGuesses = append(wrongGuesses, guess)
			life--
		}

		fmt.Println(strings.Repeat(">", 30))
	}
}

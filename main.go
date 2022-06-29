package main

import (
	"fmt"
	gp "hangman/gameplay"
	ui "hangman/user_input"
	"strings"
)

// main function
func main() {
	var guesses int
	var targetWord string
	var guessedChars []string
	var remainingWord []string

	guesses = gp.GetDifficulty()
	numLeft := guesses
	targetWord = gp.GetTargetWord()
	remainingWord = strings.Split(targetWord, "")

	fmt.Printf("You will have: %v guesses \n", guesses)
	fmt.Println(targetWord)

	for i := 0; i < guesses; i++ {
		guessedLetter := ui.GuessLetter(targetWord, guessedChars)
		guessedChars = append(guessedChars, guessedLetter)

		if strings.Contains(strings.Join(remainingWord, ""), guessedLetter) {
			fmt.Println(strings.Join(remainingWord, ""))
			guessedIndex := strings.Index(strings.Join(remainingWord, ""), guessedLetter)
			fmt.Println("you should pop, because it matched...")
			fmt.Println("here's the index: ")
			fmt.Println((guessedIndex))
			remainingWord = append(remainingWord[:guessedIndex], remainingWord[guessedIndex+1:]...)
			fmt.Println("here's the remaining word..")
			fmt.Println(remainingWord)
		}
		if len(remainingWord) == 0 {
			fmt.Println("You Won!!")
			fmt.Print("The word was: ")
			fmt.Println(targetWord)
			return
		}

		numLeft--
		fmt.Printf("You have: %v more guesses \n", numLeft)
		fmt.Println("remaining word:")
		fmt.Println(remainingWord)

		if numLeft == 0 {
			fmt.Println("GameOver!!")
			fmt.Print("The word was: ")
			fmt.Println(targetWord)
		}
	}

}

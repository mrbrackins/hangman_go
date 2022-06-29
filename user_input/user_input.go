package user_input

import (
	"fmt"
	"strings"
)

func GetLetter() string {
	var letter string
	fmt.Print("Type a letter: ")
	fmt.Scan(&letter)
	fmt.Println("Your letter is:", letter)
	return letter

}

func GuessLetter(word string, guessedChars []string) string {

	var letter string
	fmt.Print("Type a letter: ")
	fmt.Scan(&letter)
	fmt.Println("Your letter is:", letter)

	if strings.Contains(strings.Join(guessedChars, ""), letter) {
		fmt.Println("You have already guessed that letter!")

	}

	if strings.Contains(word, letter) {
		fmt.Println("good!")

	} else {
		fmt.Println("Doesn't contain the letter!")

	}
	return letter

}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

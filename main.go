package main

import (
	"fmt"
	"strings"
	ui "hangman/user_input"
	gp "hangman/gameplay"
)


// main function
func main() {
	var guesses int
	var targetWord string

	guesses = gp.GetDifficulty()
	numLeft := guesses
	targetWord = gp.GetTargetWord()
	
	fmt.Printf("You will have: %v guesses \n", guesses)
	fmt.Println(targetWord)

	for i:=0; i < guesses; i++ {
		letter := ui.GetLetter()
		if(strings.Contains(targetWord, letter)){
			fmt.Println("contains the letter!")
			numLeft--
			fmt.Printf("You have: %v more guesses \n", numLeft)
		}else{
			fmt.Println("doesn't contain the letter!")
			numLeft--
			fmt.Printf("You have: %v more guesses \n", numLeft)
		}

		if(numLeft == 0 ){
			fmt.Println("GameOver!! /n")
		}
	}   



}

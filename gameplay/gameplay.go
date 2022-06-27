package gameplay

import (
	"fmt"
	"reflect"
	"regexp"
)


var guesses int 

func GetDifficulty() int {
	var difficulty int
	
	for {
		fmt.Print("Please select a difficulty: ")
		fmt.Print("Easy: 1 ")
		fmt.Print("Medium: 2 ")
		fmt.Print("Hard: 3 ")
		fmt.Scan(&difficulty)
			if(validatedDifficulty(difficulty)){
				return setGuesses(difficulty)
			} else {
				fmt.Println("Sorry, it must be a number between 1 and 3")
				
			}
	}


	return difficulty
}

func validatedDifficulty(d int) bool{

	//check to see if input is type int
	//used Sprint to cast interface to string so I can compare
	if(fmt.Sprint(reflect.TypeOf(d)) == "int" && d <= 3){

		
		if(regexp.MustCompile(`1`).MatchString(fmt.Sprint(d)) || regexp.MustCompile(`2`).MatchString(fmt.Sprint(d)) || regexp.MustCompile(`3`).MatchString(fmt.Sprint(d))){
			fmt.Println("yo")
			return true
		}
		}
return false
}

func setGuesses(d int) int {
	switch d {
                case 1:
                    return 10
                    break
                case 2:
                    return 6
                case 3:
                    return 4
                default:
                    fmt.Println("Sorry, You have to enter 1, 2, or 3")
                   
            }
		return 0
}

func GetTargetWord() string {
	return "hangman"
}
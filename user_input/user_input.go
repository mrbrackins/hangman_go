package user_input

import "fmt"


func GetLetter() string {
	var letter string
	fmt.Print("Type a letter: ")
	fmt.Scan(&letter)
	fmt.Println("Your letter is:", letter)
	return letter

}

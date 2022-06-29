package gameplay

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"time"
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
		if validatedDifficulty(difficulty) {
			return setGuesses(difficulty)
		} else {
			fmt.Println("Sorry, it must be a number between 1 and 3")

		}
	}

	return difficulty
}

func validatedDifficulty(d int) bool {

	//check to see if input is type int
	//used Sprint to cast interface to string so I can compare
	if fmt.Sprint(reflect.TypeOf(d)) == "int" && d <= 3 {

		if regexp.MustCompile(`1`).MatchString(fmt.Sprint(d)) || regexp.MustCompile(`2`).MatchString(fmt.Sprint(d)) || regexp.MustCompile(`3`).MatchString(fmt.Sprint(d)) {

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
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	var targetWord string
	lineNumber := 0
	randNumber := r.Intn(5-1) + 1
	// open file
	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		lineNumber++
		//pick word randomly
		if lineNumber == randNumber {
			targetWord = scanner.Text()
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return targetWord
}

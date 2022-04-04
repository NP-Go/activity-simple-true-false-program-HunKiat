package main

import (
	"fmt"
	"math/rand"
	"time"

	"strconv"
)

var secretValue int

func genRandomNoAndCompare(value string) {
	//do not change this variable resultMessage, secretValue
	res := ""
	//Insert your code from here
	count := 0
	randNo := randomNumber()
	for count < 1000 { // for strconv.Itoa(randNo) != strconv.Itoa(secretValue) && count << 1000
		res = compare(randNo) //
		if res != "Well Done! Your guess is correct" {
			count = count + 1
			randNo = randomNumber()
		} else {
			break
		}
	}

	fmt.Println("The secretNo is " + strconv.Itoa(randNo) + " found aft: " + strconv.Itoa(count) + " random rounds.")
}

func compare(value int) string {
	resultMessage := ""
	if strconv.Itoa(value) < strconv.Itoa(secretValue) {
		resultMessage = "Too low, try again next time!"
	}
	if strconv.Itoa(value) > strconv.Itoa(secretValue) {
		resultMessage = "Too high"
	}
	if strconv.Itoa(value) == strconv.Itoa(secretValue) {
		resultMessage = "Well Done! Your guess is correct"
	}
	return resultMessage
}

// this creates a fixed array
// based on fixed pre-declared inputs
func makeArrayRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// produce a random int from 0 to 99
func randomNumber() int {
	arrayInt := makeArrayRange(0, 99)
	return arrayInt[rand.Intn(len(arrayInt))]
}

func main() {
	secretValue = 88
	// var guess int
	var ghostVal string
	// fmt.Println("Enter an integer value: ")
	fmt.Println("Test your luck and hack a out a secret# - hit any key to continue")
	fmt.Scanln(&ghostVal)
	genRandomNoAndCompare(ghostVal)
}

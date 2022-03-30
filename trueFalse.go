package main

import (
	"fmt"
	"math/rand"
	"time"

	"strconv"
)

// func compare(value int) string {
func compare(value string) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 88
	//Insert your code from here
	count := 0
	randNo := randomNumber()
	for strconv.Itoa(randNo) != strconv.Itoa(secretValue) && count < 10000 {
		count = count + 1
		randNo = randomNumber()
	}

	fmt.Println("The secretNo is " + strconv.Itoa(randNo) + " found aft: " + strconv.Itoa(count) + " random rounds.")
	//do not remove this line
	return resultMessge
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
	// var guess int
	var guess string
	// fmt.Println("Enter an integer value: ")
	fmt.Println("Test your luck and hack a out a secret# - hit any key to continue")
	fmt.Scanln(&guess)
	compare(guess)
}

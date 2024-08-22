package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	actualNum := rand.Intn(10000000)
	tries := 0
	for {
		guessNum := rand.Intn(10000000)
		if guess(guessNum, actualNum) == 3 {
			fmt.Printf("The number was %d\n", guessNum)
			break
		}
		tries++
	}
	fmt.Printf("It took %d tries\n", tries)

}

func guess(guessNum, actualNum int) int {
	if guessNum != actualNum {
		if guessNum > actualNum {
			return 1
		}
		if guessNum < actualNum {
			return 2
		}
	} else {
		return 3
	}
	return 0
}

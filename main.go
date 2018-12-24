package main

import (
	"fmt"
	"math/rand"
)

func emit(c chan string) {
	words := []string{"the", "quick", "brown", "fox"}

	for _, word := range words {
		//fmt.Printf("word: %s, ", word)
		c <- word
	}
	fmt.Printf("\nCompleted\n")

	//close(c)
}

func emitInt(c chan int) {
	for {
		c <- rand.Intn(1000)
	}
}

func makeID(idChan chan int) int {

	id := 0
	for {
		idChan <- id
		id++
	}

}

func makeRandoms(c chan int) {

	for { // infinite loop, passing back data
		c <- rand.Intn(1000)
	}
}

/*
	Anthonys-MacBook-Pro:go mcclayac$ godoc math/rand Intn
	func Intn(n int) int
	Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	from the default Source. It panics if n <= 0.

	type Rand struct {
		// contains filtered or unexported fields
	}

	A Rand is a source of random numbers.

	func (r *Rand) Intn(n int) int
	Intn returns, as an int, a non-negative pseudo-random number in [0,n).
	It panics if n <= 0.


*/

func main() {

	//fmt.Printf("Word Channel Example\n\n")
	//wordChannel := make(chan string)
	////randoms := make(chan int)
	//
	////randoms = randoms
	////go emit(wordChannel)
	//go emit(wordChannel)
	//fmt.Printf("\nContinues Main\n")
	//
	//for word := range wordChannel { // range over a channel,
	//	// recieve until the channel is close
	//	fmt.Printf("%s ", word)
	//}

	fmt.Printf("\n\n=================================\n\n")
	fmt.Printf("emitInt  Randoms\n\n")

	randoms := make(chan int)

	go emitInt(randoms)

	for count := 0; count < 100; count++ { // range over a channel,
		// recieve until the channel is close
		//fmt.Printf("%d ", randoms)
		fmt.Printf("%d ", randoms)

	}

	//for intRandom := range randoms {
	//	fmt.Printf("%d ", intRandom)
	//}
	fmt.Printf("\n\n=================================\n\n")
	fmt.Printf("ID Channel\n\n")

	idChan := make(chan int)

	go makeID(idChan)

	fmt.Printf("id : %d\n", <-idChan)
	fmt.Printf("id : %d\n", <-idChan)
	fmt.Printf("id : %d\n", <-idChan)
	fmt.Printf("id : %d\n", <-idChan)
	fmt.Printf("id : %d\n", <-idChan)

	//go makeRandoms(randoms)
	//
	//for randomNum := range randoms {
	//	fmt.Printf(",%d",randomNum)
	//}

	// Make ID's
	//fmt.Printf("Make ID's\n\n" )
	//idChannel := make(chan int)
	//go makeID(idChannel)
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("\trecieved ID :%d\n", <-idChannel)
	//}

	/*
		// recieve the first word
		word := <- wordChannel
		fmt.Printf("word #1 is: %s\n", word)

		// recieve the second word
		word = <- wordChannel
		fmt.Printf("word #2 is: %s\n", word)

		// recieve the third word
		word = <- wordChannel
		fmt.Printf("word #3 is: %s\n", word)

		// recieve the fourth word
		word = <- wordChannel
		fmt.Printf("word #4 is: %s\n", word)

		// channel closed
		// recieve the fifith word
		 word, okc := <- wordChannel
		fmt.Printf("word #5 is: %s\tok: %t\n", word, okc)
	*/
}

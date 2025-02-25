package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	numbersCh := make(chan int)
	squaresCh := make(chan int)

	go generateNumbers(numbersCh)
	go squareNumbers(numbersCh, squaresCh)

	var results []int
	for squared := range squaresCh {
		results = append(results, squared)
		fmt.Printf("Squared: %d\n", squared)
	}
	fmt.Println("Results: ", results)
}

func generateNumbers(ch chan int) {
	const arrLength = 10
	arr := [arrLength]int{}
	for i := 0; i < len(arr); i++ {
		number := rand.Intn(101)
		arr[i] = number
		ch <- number
	}
	fmt.Println("Initial numbers: ", arr)
	close(ch)
}

func squareNumbers(inChannel, outChannel chan int) {
	arr := make([]int, 0)
	for number := range inChannel {
		squared := int(math.Pow(float64(number), 2))
		arr = append(arr, squared)
		outChannel <- squared
	}
	fmt.Println("Squared numbers: ", arr)
	close(outChannel)
}

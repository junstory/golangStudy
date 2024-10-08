package main

// #include<stdio.h>
// void printHelloWorld(){
//     printf("Hello, World!\n");
// }
import (
	"C"
	"fmt"
)

func squareIt(inputChan, outputChan chan int) {
	for x := range inputChan {
		outputChan <- x * x
	}
}

func main() {
	C.printHelloWorld()
	inputChannel := make(chan int)
	outputChannel := make(chan int, 10)
	go squareIt(inputChannel, outputChannel)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		inputChannel <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChannel)
	}
	close(inputChannel)
}

package task5

import (
	"fmt"
	"sync"
)

func OddNumber(c chan string, i int) {
	c <- fmt.Sprintf("Receive an odd number: %v", i)
}

func EvenNumber(c chan string, i int) {
	c <- fmt.Sprintf("Receive an even number: %v", i)
}

func Consume(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range c { // Read from the channel
		fmt.Println(i)
	}
}

func OddEvenChannel() {
	var wg sync.WaitGroup
	channelOdd := make(chan string)
	channelEven := make(chan string)

	go func(){
		for i := 1; i < 21; i++ {
			if(i % 2 == 0){
				OddNumber(channelEven, i)
			}else{
				EvenNumber(channelEven, i)
			}
		}
		close(channelEven)
		close(channelOdd)
	}()

	wg.Add(2) // Add for the Consumer

	go Consume(channelOdd, &wg)
	go Consume(channelEven, &wg)

	wg.Wait()
}

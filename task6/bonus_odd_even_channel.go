package task6

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

func ErrorNumber(c chan string, i int) {
	c <- fmt.Sprintf("Error: number %v is greater than 20", i)
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
	channelError := make(chan string)

	go func(){
		for i := 1; i < 23; i++ {
			if(i < 21){
				if(i % 2 == 0){
					OddNumber(channelEven, i)
				}else{
					EvenNumber(channelEven, i)
				}
			}else{
				ErrorNumber(channelError, i)
			}
		}
		close(channelEven)
		close(channelOdd)
		close(channelError)
	}()

	wg.Add(3) // Add for the Consumer

	go Consume(channelOdd, &wg)
	go Consume(channelEven, &wg)
	go Consume(channelError, &wg)

	wg.Wait()
}

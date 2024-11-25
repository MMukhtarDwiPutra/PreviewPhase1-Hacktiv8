package task3

import (
	"fmt"
	"sync"
)

// Buat fungsi produce untuk memasukan string ke dalam channel
func Produce(c chan string) {
	for i := 1; i < 11; i++ {
		c <- fmt.Sprintf("%v", i)
	}
	close(c)
}

// Buat fungsi consume untuk receive semua isi dalam channel
func Consume(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range c { // Read from the channel
		fmt.Println(i)
	}
}

func GoChannel() {
	// Buat wait group
	var wg sync.WaitGroup

	// Buat unbuffered channel
	c := make(chan string)

	// Add 1 to counter wait group
	wg.Add(1)

	// Start the Producer function
	go Produce(c)

	// Start the Consumer function
	go Consume(c, &wg)

	wg.Wait() // Wait for the Consumer to finish
}

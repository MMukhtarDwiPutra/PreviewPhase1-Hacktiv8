package task4

import (
	"fmt"
	"sync"
)

// Buat fungsi produce untuk memasukan string ke dalam channel
func Produce(c chan string) {
	for i := 1; i < 11; i++ {
		c <- fmt.Sprintf("%v", i)
	}
	close(c) // Close the channel after all items are sent
}

// Buat fungsi consume untuk receive semua isi dalam channel
func Consume(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range c { // Read from the channel
		fmt.Println(i)
	}
}

func BufferedChannel() {
	var wg sync.WaitGroup

	// Buat buffered channel, dengan isi 5
	// Maksimal isi channel adalah 5, jika masih ada 5, maka channel harus dikeluarkan terlebih dahulu (consume)
	c := make(chan string, 5)

	wg.Add(1) // Add for the Consumer

	// Start the Producer
	go Produce(c)

	// Start the Consumer
	go Consume(c, &wg)

	wg.Wait() // Wait for the Consumer to finish
}

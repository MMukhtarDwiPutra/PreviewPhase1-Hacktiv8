package task5

import (
	"fmt"
	"sync"
)

// Buat mengisi odd number ke dalam channel odd nanti
func OddNumber(c chan string, i int) {
	c <- fmt.Sprintf("Receive an odd number: %v", i)
}

// Buat mengisi even number ke dalam channel even nanti
func EvenNumber(c chan string, i int) {
	c <- fmt.Sprintf("Receive an even number: %v", i)
}

// Buat print semua isi dalam channel nanti
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

	// 1 Go routine, berisi pengecekan odd even number
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

	// Add 2 to counter wait group
	wg.Add(2) // Add for the Consumer

	// Jalankan consume channel odd even error secara bersamaan
	go Consume(channelOdd, &wg)
	go Consume(channelEven, &wg)

	// Tunggu semua go routine selesai, menunggu wg.Done()
	wg.Wait()
}

package task6

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

// Buat mengisi error number ke dalam channel even nanti
func ErrorNumber(c chan string, i int) {
	c <- fmt.Sprintf("Error: number %v is greater than 20", i)
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
	channelError := make(chan string)

	// 1 Go routine, berisi pengecekan odd even dan error number
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

		// Close semua channel
		close(channelEven)
		close(channelOdd)
		close(channelError)
	}()

	// Add 3 count to wait group
	wg.Add(3)

	// Jalankan consume channel odd even error secara bersamaan
	go Consume(channelOdd, &wg)
	go Consume(channelEven, &wg)
	go Consume(channelError, &wg)

	// Tunggu semua go routine selesai, menunggu wg.Done()
	wg.Wait()
}

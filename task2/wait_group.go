package task2

import(
	"sync"
	"fmt"
	"time"
)

// Buat fungsi print number dengan memakai wait group agar print secara order
func PrintNumbers(wg *sync.WaitGroup){
	// Kurangin 1 count to wait group
	defer wg.Done()
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

// Buat fungsi print letter dengan memakai wait group agar print secara order
func PrintLetters(wg *sync.WaitGroup){
	// Kurangin 1 count to wait group
	defer wg.Done()
	for i := 'a'; i < 'k'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func GoWaitGroup(){
	// Buat wait group
	var wg sync.WaitGroup

	// Add 2 counter to wait group
	wg.Add(2)
	go PrintNumbers(&wg)
	go PrintLetters(&wg)

	// Wait 1 detik agar bisa melihat hasilnya
	time.Sleep(1 * time.Second)

	wg.Wait()

}
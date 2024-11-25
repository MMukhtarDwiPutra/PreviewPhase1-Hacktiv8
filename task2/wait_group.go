package task2

import(
	"sync"
	"fmt"
	"time"
)

func PrintNumbers(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

func PrintLetters(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 'a'; i < 'k'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func GoWaitGroup(){
	var wg sync.WaitGroup

	wg.Add(2)
	go PrintNumbers(&wg)
	go PrintLetters(&wg)

	time.Sleep(1 * time.Second)

	wg.Wait()

}
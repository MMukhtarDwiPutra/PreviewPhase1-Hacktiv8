package task1

import(
	"fmt"
	"time"
)

func PrintNumbers(){
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

func PrintLetters(){
	for i := 'a'; i < 'k'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func BasicGoroutine(){
	go PrintNumbers()
	go PrintLetters()

	time.Sleep(1 * time.Second)
}
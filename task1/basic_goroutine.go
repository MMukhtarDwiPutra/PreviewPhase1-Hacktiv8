package task1

import(
	"fmt"
	"time"
)

// Buat fungsi print number
func PrintNumbers(){
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

// Buat fungsi print letters
func PrintLetters(){
	for i := 'a'; i < 'k'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func BasicGoroutine(){
	// Jalankan fungsi print number dan print letters secara bersamaan (go routine)
	go PrintNumbers()
	go PrintLetters()

	time.Sleep(1 * time.Second)
}
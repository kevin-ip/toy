package main

import (
	"fmt"
	_ "github.com/KimMachineGun/automemlimit"
	_ "go.uber.org/automaxprocs"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {

	fmt.Printf("GOMAXPROCS is set to: %d\n", runtime.GOMAXPROCS(0))

	limit := debug.SetMemoryLimit(-1) // Retrieve the current limit
	fmt.Printf("GOMEMLIMIT is set to: %d bytes\n", limit)

	done := make(chan struct{})

	// Optional: Do some work in a goroutine
	go func() {
		for {
			fmt.Println("Working in the background...")
			time.Sleep(5 * time.Second)
		}
	}()

	// Block forever
	<-done
	fmt.Println("This line is never reached")
}

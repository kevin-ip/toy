package main

import (
	"fmt"
	_ "github.com/KimMachineGun/automemlimit"
	_ "go.uber.org/automaxprocs"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		for {
			fmt.Println("Working in the background...")
			time.Sleep(5 * time.Second)
		}
	}()

	// Block forever
	<-done
	fmt.Println("This line should never be reached")
}

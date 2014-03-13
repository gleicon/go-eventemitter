package main

import (
	"fmt"
	ee "github.com/gleicon/go-eventemitter"
	"os"
	"runtime/pprof"
	"time"
)

func printer(wot []byte) {
	fmt.Println(string(wot))
}

func printer2(wot []byte) {
	fmt.Println("2: ", string(wot))
}

func main() {
	myee := ee.NewEventEmitter()
	myee.On("jazz", printer)
	myee.On("jazz", printer2)
	myee.Emit("jazz", []byte("asdadsa"))
	for i := 0; i < 5; i++ {
		st := fmt.Sprintf("goroutine: %d", i)
		time.Sleep(500 * time.Millisecond)
		myee.Emit("jazz", []byte(st))
	}
	fmt.Println(len(myee.Listeners("jazz")))
	fmt.Println(pprof.WriteHeapProfile(os.Stdout))
}

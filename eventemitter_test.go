package eventemitter

import (
	"fmt"
	"testing"
)

func TestEventEmitter(t *testing.T) {
	myee := NewEventEmitter()
	myee.On("jazz", func(wot []byte) {
		if string(wot) != "el test" {
			t.Error("Invalid value")
		}
	})
	myee.Emit("jazz", []byte("el test"))
}

func TestEventEmitterListeners(t *testing.T) {
	myee := NewEventEmitter()
	myee.On("jazz", func(wot []byte) { fmt.Println(wot) })
	myee.On("jazz", func(wot []byte) { fmt.Println(wot) })
	if len(myee.Listeners("jazz")) != 2 {
		t.Error("Invalid listeners value")
	}
}

func TestRemoveAllListeners(t *testing.T) {
	myee := NewEventEmitter()
	myee.On("jazz", func(wot []byte) { fmt.Println(wot) })
	myee.On("jazz", func(wot []byte) { fmt.Println(wot) })
	myee.RemoveAllListeners("jazz")
	if len(myee.Listeners("jazz")) != 0 {
		t.Error("Invalid listeners value")
	}
	myee.On("jazz", func(wot []byte) {
		if string(wot) != "el test" {
			t.Error("Invalid value")
		}
	})
	myee.Emit("jazz", []byte("el test"))
}

func BenchmarkEventEmitterSequential(b *testing.B) {
	myee := NewEventEmitter()
	myee.On("jazz", func(wot []byte) { fmt.Println(string(wot)) })
	for i := 0; i < 5; i++ {
		st := fmt.Sprintf("iteration: %d", i)
		myee.Emit("jazz", []byte(st))
	}
}

func BenchmarkEventEmitterGoroutines(b *testing.B) {
	myee := NewEventEmitter()
	myee.On("jazz", func(wot []byte) { fmt.Println(string(wot)) })
	for i := 0; i < 5; i++ {
		go func(idx int) {
			st := fmt.Sprintf("goroutine: %d", idx)
			myee.Emit("jazz", []byte(st))
		}(i)
	}
}

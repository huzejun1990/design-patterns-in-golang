// @Author huzejun 2024/4/28 19:43:00
package Observer

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data int
}

type Observer interface {
	NotifyCallback(event Event)
}

type Subject interface {
	AddListener(observer Observer)
	RemoverListener(observer Observer)
	Notify(event Event)
}

type eventObserver struct {
	ID   int
	Time time.Time
}

type eventSubject struct {
	Observers sync.Map
}

func (e eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Recived:%d after %v\n", event.Data, time.Since(e.Time))
}

func (e *eventSubject) AddListener(obs Observer) {
	e.Observers.Store(obs, struct{}{})
}

func (e *eventSubject) RemoverListener(obs Observer) {
	e.Observers.Delete(obs)
}

func (e *eventSubject) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		//for i, j := 0, 1; i < 0; i, j = i+j, i {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

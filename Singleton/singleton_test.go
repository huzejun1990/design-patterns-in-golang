// @Author huzejun 2024/4/27 21:43:00
package Singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGEtInstance(t *testing.T) {
	/*	p := GEtInstance()
		p.IncrementAge()*/
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			IncrementAge()
		}()
		go func() {
			defer wg.Done()
			IncrementAge2()
		}()
	}
	wg.Wait()
	age := p.GetAge()
	fmt.Println(age)
}

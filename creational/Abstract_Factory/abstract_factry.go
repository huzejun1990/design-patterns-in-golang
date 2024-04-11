// @Author huzejun 2024/4/11 20:17:00
package Abstract_Factory

import "fmt"

type Lunch interface {
	Cook()
}

type Rise struct {
}

func (r *Rise) Cook() {
	fmt.Println("it is rise")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("it is Tomato.")
}

type LuchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {
}

func NewSimpleLunchFactory() LuchFactory {
	return &SimpleLunchFactory{}
}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

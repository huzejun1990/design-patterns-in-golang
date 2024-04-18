// @Author huzejun 2024/4/18 17:00:00
package template

import "fmt"

type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}

type Worker struct {
	WorkInterface
}

func NewWorker(w WorkInterface) *Worker {
	return &Worker{w}
}

func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Corder struct {
}

func (c *Corder) GetUp() {
	fmt.Println("coder GetUp.")
}

func (c *Corder) Work() {
	fmt.Println("coder Coding.")
}

func (c *Corder) Sleep() {
	fmt.Println("coder Sleep.")
}

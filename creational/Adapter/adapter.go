// @Author huzejun 2024/4/28 17:54:00
package Adapter

import "fmt"

type Target interface {
	Execute()
}

type Adaptee struct {
}

func (a *Adaptee) SpecificExecute() {
	fmt.Println("充电...")
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Execute() {
	a.SpecificExecute()
}

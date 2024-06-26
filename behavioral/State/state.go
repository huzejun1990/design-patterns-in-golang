// @Author huzejun 2024/4/20 17:53:00
package State

import "fmt"

type State interface {
	On(m *Machine)
	Off(m *Machine)
}

type Machine struct {
	current State
}

func NewMachine() *Machine {
	return &Machine{NewOFF()}
}

func (m *Machine) setCurrent(s State) {
	m.current = s
}

func (m *Machine) On() {
	m.current.On(m)
}

func (m *Machine) Off() {
	m.current.Off(m)
}

type ON struct {
}

func NewON() State {
	return &ON{}
}

func (o *ON) On(m *Machine) {
	fmt.Println("设备已经开启。。。")
}

func (o *ON) Off(m *Machine) {
	fmt.Println("从On的状态到Off")
	m.setCurrent(NewON())
}

type OFF struct {
}

func NewOFF() State {
	return &OFF{}
}

func (o *OFF) On(m *Machine) {
	fmt.Println("从Off状态到On")
	m.setCurrent(NewON())
}

func (o *OFF) Off(m *Machine) {
	fmt.Println("已经关闭")
}

// @Author huzejun 2024/4/21 18:46:00
package State

import "testing"

func TestState(t *testing.T) {
	machine := NewMachine()
	machine.Off()
	machine.On()
	machine.On()
	machine.Off()

}

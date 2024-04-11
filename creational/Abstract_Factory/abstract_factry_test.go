// @Author huzejun 2024/4/11 22:37:00
package Abstract_Factory

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	factory := NewSimpleLunchFactory()
	food := factory.CreateFood()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}

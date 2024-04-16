// @Author huzejun 2024/4/17 3:21:00
package Builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder()
	dircotor := NewDirector(&builder)
	dircotor.Construct()
	result := builder.GetResult()
	fmt.Println(result.Built)
}

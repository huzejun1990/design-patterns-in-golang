// @Author huzejun 2024/4/23 19:16:00
package Interpreter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterpret(t *testing.T) {
	expression := "w x z +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["w"] = &Integer{6}
	variables["x"] = &Integer{10}
	variables["z"] = &Integer{41}
	result := sentence.Interpret(variables)
	assert.Equal(t, 51, result)

}

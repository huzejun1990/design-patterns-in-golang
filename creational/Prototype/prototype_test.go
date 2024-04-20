// @Author huzejun 2024/4/20 17:43:00
package Prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "出来玩呀"
	p := ConcretePrototype{name: name}
	newProto := p.Clone()
	assert.Equal(t, name, newProto.Name())
}

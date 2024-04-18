// @Author huzejun 2024/4/18 19:40:00
package Visitor

import "testing"

func TestElement_Accept(t *testing.T) {
	e := new(Element)
	e.Accept(new(WeiBoVisitor))
	e.Accept(new(IQIYIVisitor))
}

// @Author huzejun 2024/4/18 19:36:00
package Visitor

import "fmt"

type IVisitor interface {
	Visitor()
}

type WeiBoVisitor struct {
}

func (w WeiBoVisitor) Visitor() {
	fmt.Println("Visitor WeiBo")
}

type IQIYIVisitor struct {
}

func (i IQIYIVisitor) Visitor() {
	fmt.Println("Visit IQiYi")
}

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct {
}

func (e Element) Accept(v IVisitor) {
	v.Visitor()
}

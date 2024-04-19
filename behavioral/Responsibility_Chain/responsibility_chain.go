// @Author huzejun 2024/4/19 17:14:00
package Responsibility_Chain

import "strconv"

type Handler interface {
	Handler(handlerID int) string
}

type handler struct {
	name      string
	next      Handler
	handlerID int
}

func NewHandler(name string, next Handler, handlerID int) *handler {
	return &handler{
		name:      name,
		next:      next,
		handlerID: handlerID,
	}
}

func (h *handler) Handler(handlerID int) string {
	if h.handlerID == handlerID {
		return h.name + " handled " + strconv.Itoa(handlerID)
	}
	return h.next.Handler(handlerID)
}

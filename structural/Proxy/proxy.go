// @Author huzejun 2024/4/28 18:12:00
package Proxy

import (
	"fmt"
	"strconv"
)

type ITask interface {
	RentHouse(desc string, price int)
}

type Task struct {
}

func (t *Task) RentHouse(desc string, price int) {
	fmt.Println(fmt.Sprintf("租房子的地址%s,价钱%s,中介费%s,", desc, strconv.Itoa(price), strconv.Itoa(price)))
}

type AgentTask struct {
	task *Task
}

func NewAgentTask() *AgentTask {
	return &AgentTask{task: &Task{}}

}

func (t *AgentTask) RentHouse(desc string, price int) {
	t.task.RentHouse(desc, price)
}

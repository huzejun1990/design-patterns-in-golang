// @Author huzejun 2024/4/28 18:30:00
package Proxy

import "testing"

func TestAgentTask_RentHouse(t *testing.T) {
	agent := NewAgentTask()
	agent.RentHouse("北京", 8000)
}

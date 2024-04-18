// @Author huzejun 2024/4/18 17:30:00
package template

import "testing"

func TestWorker_Daily(t *testing.T) {
	worker := NewWorker(&Corder{})

	worker.Daily()
}

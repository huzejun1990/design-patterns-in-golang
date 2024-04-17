// @Author huzejun 2024/4/17 21:12:00
package Command

import "testing"

func TestCommand_Execute(t *testing.T) {
	laowang := NewPerson("wang", NewCommand(nil, nil))
	laozhang := NewPerson("zhang", NewCommand(&laowang, laowang.Listen))
	laofeng := NewPerson("feng", NewCommand(&laozhang, laozhang.Buy))
	laoding := NewPerson("ding", NewCommand(&laofeng, laofeng.Cook))
	laoli := NewPerson("li", NewCommand(&laoding, laoding.Wash))

	laoli.Talk()
}

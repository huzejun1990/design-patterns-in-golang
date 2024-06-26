// @Author huzejun 2024/4/17 4:46:00
package Bridage

import "testing"

func TestCircle_Draw(t *testing.T) {
	red := Circle{}
	red.Constructor(100, 100, 10, &RedCircle{})
	red.Draw()

	yellow := Circle{}
	yellow.Constructor(200, 200, 20, &YellowCircle{})
	yellow.Draw()
}

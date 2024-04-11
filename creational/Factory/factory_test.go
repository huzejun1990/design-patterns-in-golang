// @Author huzejun 2024/4/11 19:54:00
package Factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("d").GetFood()
	NewRestaurant("q").GetFood()
}

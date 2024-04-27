// @Author huzejun 2024/4/27 22:56:00
package Singleton

func IncrementAge() {
	p := GEtInstance()
	p.IncrementAge()
}

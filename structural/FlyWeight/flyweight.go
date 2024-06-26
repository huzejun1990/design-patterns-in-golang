// @Author huzejun 2024/4/21 22:08:00
package FlyWeight

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name: name}
}

type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}

func NewFlyWeiFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]*FlyWeight)}
}

func (f *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight {
	weight, ok := f.pool[name]
	if !ok {
		weight = NewFlyWeight(name)
		f.pool[name] = weight
	}
	return weight
}

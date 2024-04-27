// @Author huzejun 2024/4/25 19:49:00
package Singleton

import "sync"

var (
	p    *Pet
	once sync.Once
)

func init() {
	once.Do(func() {
		p = &Pet{}
	})
}

func GEtInstance() *Pet {
	return p
}

type Pet struct {
	name string
	age  int
	mux  sync.Mutex
}

func (s *Pet) SetName(n string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.name = n
}

func (p *Pet) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *Pet) IncrementAge() {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.age++
}

func (p *Pet) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}

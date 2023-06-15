package pool

import (
	"fmt"
	"time"
)

type Store interface {
	SaveName(string)
	GetName(string) string
}

type Pool struct {
	store Store
}

func NewPool(store Store) (*Pool, error) {
	fmt.Println("abc")

	return &Pool{store: store}, nil
}

func (p *Pool) run() {
	for {
		p.store.SaveName("abc")
		fmt.Println(p.store.GetName("abc"))
		time.Sleep(5 * time.Second)
	}
}

func (p *Pool) Start() {
	go p.run()
}

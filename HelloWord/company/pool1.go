package main

import (
	"fmt"
	"sync"
)

type Pool struct {
	mutex   sync.Mutex
	objects []interface{}
	new     New
}

// New 是定义了一个叫做New2的函数类型，返回一个interface
type New func() interface{}

func NewPool(new New) *Pool {
	return &Pool{
		objects: make([]interface{}, 0),
		new:     new,
	}
}

func (p *Pool) Get() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.objects) > 0 {
		obj := p.objects[0]
		p.objects = p.objects[1:]
		return obj
	} else {
		return p.new()
	}
}

func (p *Pool) Put(obj interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.objects = append(p.objects, obj)
}

func main() {
	pool := NewPool(func() interface{} {
		fmt.Println("new")
		return 1
	})
	fmt.Println(pool)
}

//9:05

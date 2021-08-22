package main

import "time"

type GLimit struct {
	Limit int
	ch    chan struct{}
}

func NewGLimit(limit int) *GLimit {
	return &GLimit{
		Limit: limit,
		ch:    make(chan struct{}, limit),
	}
}

func (glimit *GLimit) Run(foo func()) {
	glimit.ch <- struct{}{}
	go func() {
		foo1()
		<-glimit.ch
	}()
}

func add() {
	time.Sleep(100 * time.Millisecond)
	_ = 4
}

func main() {
	gl := NewGLimit(100)
	for {
		gl.Run(add)
	}
}

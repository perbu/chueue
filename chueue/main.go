package chueue

import "sync"

type Chueue struct {
	mu sync.Mutex
	ch chan struct{}
}

func NewChueue() *Chueue {
	c := &Chueue{
		ch: make(chan struct{}),
	}
	return c
}

func (c *Chueue) Release() {
	c.mu.Lock()
	defer c.mu.Unlock()
	close(c.ch)
	c.ch = make(chan struct{})
}

func (c *Chueue) Wait() {
	c.mu.Lock()
	ch := c.ch
	c.mu.Unlock()
	if ch == nil {
		return
	}
	<-ch
}

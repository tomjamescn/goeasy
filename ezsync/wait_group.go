package ezsync

import "sync"

type WaitGroupLimit struct {
	*sync.WaitGroup
	ch chan struct{}
}

func NewWaitGroupLimit(limit int) *WaitGroupLimit {
	if limit <= 0 {
		panic("limit must greater than 0")
	}
	return &WaitGroupLimit{
		WaitGroup: &sync.WaitGroup{},
		ch:        make(chan struct{}, limit),
	}
}

func (w *WaitGroupLimit) Add(delta int) {
	w.WaitGroup.Add(delta)
	w.ch <- struct{}{}
}

func (w *WaitGroupLimit) Done() {
	w.WaitGroup.Done()
	<-w.ch
}

func (w *WaitGroupLimit) Wait() {
	w.WaitGroup.Wait()
}

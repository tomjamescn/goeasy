package sync

import gosync "sync"

type WaitGroupLimit struct {
	*gosync.WaitGroup
	ch chan struct{}
}

func NewWaitGroupLimit(limit int) *WaitGroupLimit {
	if limit <= 0 {
		panic("limit must greater than 0")
	}
	return &WaitGroupLimit{
		WaitGroup: &gosync.WaitGroup{},
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

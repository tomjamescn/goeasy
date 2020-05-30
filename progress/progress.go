package progress

import (
	"sync"

	tm "github.com/buger/goterm"
)

type Indicator struct {
	*sync.Mutex
	seq      []string
	progress map[string]string
}

func NewIndicator() *Indicator {
	return &Indicator{
		Mutex:    &sync.Mutex{},
		progress: map[string]string{},
	}
}

func (i *Indicator) Println(a ...interface{}) {
	tm.Println(a...)
	tm.Flush()
}

func (i *Indicator) Printf(format string, a ...interface{}) {
	tm.Printf(format, a...)
	tm.Flush()
}

func (i *Indicator) Update(name string, progress string) {
	i.Lock()
	defer i.Unlock()

	if _, ok := i.progress[name]; !ok {
		i.seq = append(i.seq, name)
	}

	i.progress[name] = progress

	h := len(i.progress)

	for _, k := range i.seq {
		tm.Printf("%s %s\n", k, i.progress[k])
	}
	tm.Flush()

	tm.MoveCursorUp(h)
}

func (i *Indicator) Done() {
	i.Lock()
	defer i.Unlock()

	h := len(i.progress)

	tm.MoveCursorDown(h)
}

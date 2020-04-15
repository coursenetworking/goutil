package tmap

import (
	"sync"
	"time"
)

type item struct {
	v string
	t int64
}

type TMap struct {
	m map[string]*item
	i int64
	l sync.Mutex
}

func New(size int, live int) (m *TMap) {
	m = &TMap{m: make(map[string]*item, size), i: int64(live)}
	go func() {
		for n := range time.Tick(time.Second) {
			m.l.Lock()
			for k, v := range m.m {
				if v.t-n.Unix() < m.i {
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
	return
}

func (m *TMap) Len() int {
	return len(m.m)
}

// t: time.Now().Unix()
func (m *TMap) Put(k, v string, t int64) {
	m.l.Lock()
	m.m[k] = &item{v: v, t: t}
	m.l.Unlock()
}

func (m *TMap) Get(k string) (v string) {
	m.l.Lock()
	if it, ok := m.m[k]; ok {
		if it.t-time.Now().Unix() > m.i {
			v = it.v
		}
	}
	m.l.Unlock()

	return
}

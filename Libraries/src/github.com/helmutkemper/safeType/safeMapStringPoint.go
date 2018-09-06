package safeType

import "sync"

type SafeMapStringPoint struct {
	l    sync.Mutex
	m    map[string][2]float64
	init bool
}

func (el *SafeMapStringPoint) Clear() {
	el.init = true
	el.m = make(map[string][2]float64, 0)
}
func (el *SafeMapStringPoint) Set(k string, v [2]float64) {
	defer el.l.Unlock()

	el.l.Lock()

	if el.init == false {
		el.init = true
		el.m = make(map[string][2]float64)
	}

	el.m[k] = v
}
func (el *SafeMapStringPoint) Len() int {
	defer el.l.Unlock()

	el.l.Lock()
	return len(el.m)
}
func (el *SafeMapStringPoint) GetKey(k string) [2]float64 {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m[k]
}
func (el *SafeMapStringPoint) DeleteKey(k string) {
	defer el.l.Unlock()

	el.l.Lock()
	delete(el.m, k)
}
func (el *SafeMapStringPoint) Get() map[string][2]float64 {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m
}

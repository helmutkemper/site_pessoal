package safeType

import "sync"

type SafeMapInt64Point struct {
	l    sync.Mutex
	m    map[int64][2]float64
	init bool
}

func (el *SafeMapInt64Point) Clear() {
	el.init = true
	el.m = make(map[int64][2]float64, 0)
}
func (el *SafeMapInt64Point) Set(k int64, v [2]float64) {
	defer el.l.Unlock()

	el.l.Lock()

	if el.init == false {
		el.init = true
		el.m = make(map[int64][2]float64)
	}

	el.m[k] = v
}
func (el *SafeMapInt64Point) Len() int {
	defer el.l.Unlock()

	el.l.Lock()
	return len(el.m)
}
func (el *SafeMapInt64Point) GetKey(k int64) [2]float64 {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m[k]
}
func (el *SafeMapInt64Point) DeleteKey(k int64) {
	defer el.l.Unlock()

	el.l.Lock()
	delete(el.m, k)
}
func (el *SafeMapInt64Point) Get() map[int64][2]float64 {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m
}

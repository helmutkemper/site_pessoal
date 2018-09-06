package safeType

import "sync"

type SafeMapStringString struct {
	l    sync.Mutex
	m    map[string]string
	init bool
}

func (el *SafeMapStringString) Clear() {
	el.init = true
	el.m = make(map[string]string)
}
func (el *SafeMapStringString) Set(k string, v string) {
	defer el.l.Unlock()

	el.l.Lock()

	if el.init == false {
		el.init = true
		el.m = make(map[string]string)
	}

	el.m[k] = v
}
func (el *SafeMapStringString) Len() int {
	defer el.l.Unlock()

	el.l.Lock()
	return len(el.m)
}
func (el *SafeMapStringString) GetKey(k string) string {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m[k]
}
func (el *SafeMapStringString) DeleteKey(k string) {
	defer el.l.Unlock()

	el.l.Lock()
	delete(el.m, k)
}
func (el *SafeMapStringString) Get() map[string]string {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m
}

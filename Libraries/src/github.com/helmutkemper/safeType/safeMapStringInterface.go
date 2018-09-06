package safeType

import "sync"

type SafeMapStringInterface struct {
	l    sync.Mutex
	m    map[string]interface{}
	init bool
}

func (el *SafeMapStringInterface) Clear() {
	el.init = true
	el.m = make(map[string]interface{})
}
func (el *SafeMapStringInterface) Set(k string, v interface{}) {
	defer el.l.Unlock()

	el.l.Lock()

	if el.init == false {
		el.init = true
		el.m = make(map[string]interface{})
	}

	el.m[k] = v
}
func (el *SafeMapStringInterface) Len() int {
	defer el.l.Unlock()

	el.l.Lock()
	return len(el.m)
}
func (el *SafeMapStringInterface) GetKey(k string) interface{} {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m[k]
}
func (el *SafeMapStringInterface) DeleteKey(k string) {
	defer el.l.Unlock()

	el.l.Lock()
	delete(el.m, k)
}
func (el *SafeMapStringInterface) Get() map[string]interface{} {
	defer el.l.Unlock()

	el.l.Lock()
	return el.m
}

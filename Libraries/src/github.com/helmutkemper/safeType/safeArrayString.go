package safeType

import "sync"

type SafeArrayString struct {
	l    sync.Mutex
	a    []string
	init bool
}

func (el *SafeArrayString) Clear() {
	defer el.l.Unlock()

	el.l.Lock()
	el.init = true
	el.a = make([]string, 0)
}
func (el *SafeArrayString) Append(v string) {
	defer el.l.Unlock()

	el.l.Lock()
	if el.init == false {
		el.init = true
		el.a = make([]string, 0)
	}

	el.a = append(el.a, v)
}
func (el *SafeArrayString) Len() int {
	defer el.l.Unlock()

	el.l.Lock()
	return len(el.a)
}
func (el *SafeArrayString) GetKey(i int) string {
	defer el.l.Unlock()

	el.l.Lock()
	return el.a[i]
}
func (el *SafeArrayString) Get() []string {
	defer el.l.Unlock()

	el.l.Lock()
	if el.init == false {
		el.init = true
		el.a = make([]string, 0)
	}

	return el.a
}

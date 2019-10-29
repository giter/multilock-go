package multilock

import (
	"strings"
	"sync"
)

var mm = sync.Map{}

func New(id ...string) *sync.Mutex {

	m, _ := mm.LoadOrStore(strings.Join(id, "ᚋ"), &sync.Mutex{})
	return m.(*sync.Mutex)
}

//go:build go1.24

package sync

import (
	"reflect"
)

func (l *Mutex) IsLocked() bool {
	mu := reflect.ValueOf(&l.lock).Elem().FieldByName("mu")
	state := mu.FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}

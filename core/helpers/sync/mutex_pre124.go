//go:build !go1.24

package sync

import (
	"reflect"
)

func (l *Mutex) IsLocked() bool {
	state := reflect.ValueOf(&l.lock).Elem().FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}

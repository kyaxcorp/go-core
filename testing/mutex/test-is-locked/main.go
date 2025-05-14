package main

import "github.com/kyaxcorp/go-core/core/helpers/sync"

var m sync.Mutex

func main() {
	m.TryLock()
}

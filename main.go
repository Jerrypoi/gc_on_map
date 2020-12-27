package main

import (
	log "github.com/sirupsen/logrus"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

var (
	dataMap       = make(map[int64][]byte)
	lock          = sync.Mutex{}
	index   int64 = 0
)

func main() {
	log.Println(debug.SetGCPercent(50))

	pool := sync.Pool{New: func() interface{} {
		return new([1024]byte)
	}}

	go func() {
		ticker := time.NewTicker(time.Second * 20)
		for {
			select {
			case <-ticker.C:
				lock.Lock()
				index = 0
				log.Println("Trigger ticker")
				runtime.GC()
				debug.FreeOSMemory()
				lock.Unlock()
			}
		}
	}()

	for {
		lock.Lock()
		dataMap[index] = pool.Get().([]byte)
		index++
		lock.Unlock()
	}
}

/*func FreePtr(p unsafe.Pointer)

func Free(v interface {}) {
	FreePtr(unsafe.Pointer(reflect.ValueOf(v).Elem().Pointer()))
}*/

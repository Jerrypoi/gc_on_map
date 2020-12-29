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
	go func() {
		ticker := time.NewTicker(time.Second * 20)
		for {
			select {
			case <-ticker.C:
				lock.Lock()
				index = 0
				for k, _ := range dataMap {
					delete(dataMap, k)
				}
				dataMap = nil
				dataMap = make(map[int64][]byte)
				log.Println("Trigger ticker")
				runtime.GC()
				debug.FreeOSMemory()
				lock.Unlock()
			}
		}
	}()

	for {
		lock.Lock()
		dataMap[index] = make([]byte, 1024)
		index++
		if index%100000 == 0 {
			log.Println(index)
		}
		lock.Unlock()
	}
}

/*func FreePtr(p unsafe.Pointer)

func Free(v interface {}) {
	FreePtr(unsafe.Pointer(reflect.ValueOf(v).Elem().Pointer()))
}*/

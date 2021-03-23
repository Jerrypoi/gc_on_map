package main

import (
	log "github.com/sirupsen/logrus"
)

var (
	dataMap       = make(map[int64][]byte)
	index   int64 = 0
)

func main() {
	for {
		dataMap[index] = make([]byte, 1024)
		index++
		if index%100000 == 0 {
			log.Println(index)
		}
	}
}

/*func FreePtr(p unsafe.Pointer)

func Free(v interface {}) {
	FreePtr(unsafe.Pointer(reflect.ValueOf(v).Elem().Pointer()))
}*/

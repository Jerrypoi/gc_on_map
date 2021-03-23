package main

import (
	log "github.com/sirupsen/logrus"
)

var (
	dataArray       = make([][]byte, 0)
	index     int64 = 0
)

func main() {
	for {
		//dataMap[index] = make([]byte, 1024)
		dataArray = append(dataArray, make([]byte, 1024*1024))
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

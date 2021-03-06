package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.RWMutex

type metadata struct {
	a map[string]interface{}
}

func (m metadata) read(stopSignal chan bool) {
	for {
		select {
		case <- stopSignal:
            fmt.Println("stopping routine")
			return
		default:
			mutex.RLock()
			x := m.a["helloworld"].(int)
			x += 1
			mutex.RUnlock()
		}
	}

}

func (m metadata) write(stopSignal chan bool) {
	for true {
		select {
		case <- stopSignal:
			fmt.Println("stopping routine")
			return
		default:
			mutex.RLock()
			x := m.a["helloworld"].(int)
			mutex.RUnlock()
			mutex.Lock()
			m.a["helloworld"] = x + 1
			mutex.Unlock()
		}
	}
}

func main() {
	data := map[string]interface{}{
		"helloworld": 0,
	}
	var m metadata
	m.a = data
	fmt.Println("for detecting race")
	stopSignal := make(chan bool)
	fmt.Println("spawning read routine")
	go m.read(stopSignal)
	fmt.Println("spawning write routine")
	go m.write(stopSignal)
	startTime := time.Now()
	fmt.Printf("time now: ")
	fmt.Println(startTime)
	for true {
		timeNow := time.Now()
		timeElapsed := int(timeNow.Sub(startTime).Seconds())
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(timeElapsed)
		if timeElapsed > 4 {
            fmt.Println("stopping first routine")
			stopSignal <- true
            fmt.Println("stopping second routine")
			stopSignal <- true
			break
		}
	}
	fmt.Printf("time now: ")
	fmt.Println(time.Now())
}

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
		case <-stopSignal:
			fmt.Println("stopping routine")
			break
		default:
			mutex.RLock()
			x := m.a["helloworld"].(int)
			x += 1
			mutex.RUnlock()
		}
	}

}

func (m metadata) write(stopSignal chan bool) {
	for {
		select {
		case <-stopSignal:
			fmt.Println("stopping routine")
			break
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
	var stopSignal chan bool
	fmt.Println("spawning read routine")
	go m.read(stopSignal)
	fmt.Println("spawning write routine")
	go m.write(stopSignal)
	startTime := time.Now()
	fmt.Printf("time now: ")
	fmt.Println(startTime)
	for true {
		timeNow := time.Now()
		timeElapsed := int(timeNow.Sub(startTime).Minutes())
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(timeElapsed)
		if timeElapsed > 2 {

			stopSignal <- true
			stopSignal <- true
			break
		}
	}
	fmt.Printf("time now: ")
	fmt.Println(time.Now())
}

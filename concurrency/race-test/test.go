package main

import (
	"fmt"
	"time"
)

type metadata struct {
	a map[string]interface{}
}

func (m metadata) read(stopSignal chan bool) {
	for {
		select{
			case <- stopSignal:
				break
			default:
				x := m.a["helloworld"].(int)
				x += 1
		}
	}
	
}

func (m metadata) write(stopSignal chan bool) {
	for {
		select{
			case <- stopSignal:
				break
			default:
				fmt.Println(m.a["helloworld"].(int))
				m.a["helloworld"] = m.a["helloworld"].(int) + 1
		}
	}
}

func main() {
	data := map[string]interface{}{
		"helloworld": 0,
	}
	var m metadata;
	m.a = data
	fmt.Println("for detecting race")
	var stopSignal chan bool
	fmt.Println("spawning read routine")
	go m.read(stopSignal)
	fmt.Println("spawning write routine")
	go m.write(stopSignal)
	startTime := time.Now()
	for true {
		timeNow := time.Now()
		timeElapsed := int(timeNow.Sub(startTime).Minutes())
		if timeElapsed > 5 {
			
			stopSignal <- true
			stopSignal <- true
			break
		}
	}
}


package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func initializeRand() {
	rand.Seed(time.Now().UnixNano())
	// rand.Seed(time.Date(
	// 1995, 10, 02, 00, 00, 00, 651387237, time.UTC).UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getHit(keySize int, key string, randomize bool) {
	var resp *http.Response
	fmt.Println(key)
	if randomize == false {
		resp, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/get/?key=%s", "ilapahsi"))
	} else {
		resp, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/get/?key=%s", key))
	}
	// if err != nil {
	// 	fmt.Println(resp)
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	f, err := os.Create("error.log")
	// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bodyString := string(bodyBytes)
	// 	f.WriteString(bodyString)
	// } else {
	// 	fmt.Println(resp)
	// }

	fmt.Println(resp)

	// if resp.StatusCode == http.StatusOK {
	// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bodyString := string(bodyBytes)
	// 	fmt.Println(bodyString)
	// }
}

func putHit(samples int, keySize int, valueSize int, randomize bool, keys chan interface{}) {
	var key string
	var value string
	// var resp *http.Response
	for i := 0; i < samples; i++ {
		if randomize == false {
			_, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/put/?key=%s&value=%s", "ilapahsi", "alice"))
		} else {
			key = randStringRunes(keySize)
			value = randStringRunes(valueSize)
			_, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/put/?key=%s&value=%s", key, value))
		}

		// fmt.Println(resp)
		// if resp.StatusCode != http.StatusOK {
		// 	f, err := os.Create("error.log")
		// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	bodyString := string(bodyBytes)
		// 	f.WriteString(bodyString)
		// } else {
		// 	keys <- key
		// }

		keys <- key
		// if resp.StatusCode == http.StatusOK {
		// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	bodyString := string(bodyBytes)
		// 	fmt.Println(bodyString)
		// }
	}
}

func main() {
	randomize := flag.Bool("randomize", false, "randomize put and get")
	// fmt.Println(*randomize)
	samples := 1000
	keySize := 10
	valueSize := 20
	// concurrency := 50
	start := time.Now()
	initializeRand()
	keys := make(chan interface{}, samples)
	for i := 0; i < samples; i++ {
		go putHit(samples, keySize, valueSize, *randomize, keys)
	}
	fmt.Println("put done")
	var keyList []string
	for i := 0; i < samples; i++ {
		key := <-keys
		keyList = append(keyList, key.(string))
	}
	for i := 0; i < samples; i++ {
		go getHit(keySize, keyList[i], *randomize)
	}
	fmt.Println("get done")
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}

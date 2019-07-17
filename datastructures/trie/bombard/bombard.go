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

func getHit(keySize int, keyList []string, randomize bool) {
	var key string
	var resp *http.Response
	for i := 0; i < len(keyList); i++ {
		key = keyList[i]
		fmt.Println(key)
		if randomize == false {
			resp, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/get/?key=%s", "ilapahsi"))
		} else {
			resp, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/get/?key=%s", key))
		}
		// if err != nil {
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
}

func putHit(samples int, keySize int, valueSize int, randomize bool) []string {
	var key string
	var value string
	var resp *http.Response
	keyList := make([]string, samples)
	for i := 0; i < samples; i++ {
		key = randStringRunes(keySize)
		keyList = append(keyList, key)
		value = randStringRunes(valueSize)
		if randomize == false {
			resp, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/put/?key=%s&value=%s", "ilapahsi", "alice"))
		} else {
			resp, _ = http.Get(fmt.Sprintf("http://127.0.0.1:1337/put/?key=%s&value=%s", key, value))
		}
		// if err != nil {
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
	return keyList
}

func main() {
	randomize := flag.Bool("randomize", false, "randomize put and get")
	samples := 1000
	keySize := 10
	valueSize := 20
	// concurrency := 50
	start := time.Now()
	initializeRand()
	keyList := putHit(samples, keySize, valueSize, *randomize)
	initializeRand()
	getHit(keySize, keyList, *randomize)
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}

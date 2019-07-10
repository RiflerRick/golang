package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func initialize_rand() {
	rand.Seed(time.Date(
		1995, 10, 02, 00, 00, 00, 651387237, time.UTC).UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func get_hit(samples int, keySize int) {
	var key string
	for i := 0; i < samples; i++ {
		key = RandStringRunes(keySize)
		resp, _ := http.Get(fmt.Sprintf("http://127.0.0.1:1337/get/%s", key))
		// if err != nil {
		// 	fmt.Println(resp)
		// }
		fmt.Println(resp)
	}
}

func put_hit(samples int, keySize int, valueSize int) {
	var key string
	var value string
	for i := 0; i < samples; i++ {
		key = RandStringRunes(keySize)
		value = RandStringRunes(valueSize)
		resp, _ := http.Get(fmt.Sprintf("http://127.0.0.1:1337/put/%s/%s", key, value))
		// if err != nil {
		// 	fmt.Println(resp)
		// }
		fmt.Println(resp)
	}
}

func main() {
	samples := 10000
	keySize := 10
	valueSize := 20
	// concurrency := 50
	start := time.Now()
	initialize_rand()
	put_hit(samples, keySize, valueSize)
	initialize_rand()
	get_hit(samples, keySize)
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}

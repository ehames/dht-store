package dhtstore

import "fmt"

type Entry struct {
	Key   string `json:key`
	Value string `json:"value"`
}

func Put(key string, value string) {
	fmt.Println(key + " => " + value)
}

func Get(key string) string {
	return ""
}

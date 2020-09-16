package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	for {
		err := ioutil.WriteFile("/static/exchange", []byte(time.Now().String()), 0644)
		if err != nil {
			fmt.Println("Error writing to file")
		}
		time.Sleep(time.Second * 1)
	}
}

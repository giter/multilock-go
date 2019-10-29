package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/giter/multilock-go"
)

func main() {

	var line int

	a := func(i int) {

		lock := multilock.New("a", strconv.Itoa(i%20))

		lock.Lock()
		defer lock.Unlock()

		line++
		fmt.Println("L", line, "->", i)
	}

	for i := 0; i < 1000; i++ {
		go a(i)
	}

	time.Sleep(time.Second)
}

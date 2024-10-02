package main

import (
	"log"
	"time"

	"go-panic-recover/src/util"
)

func doPanic() {
	defer util.CatchError()

	//? Test panic
	var data any = "abcdef"
	log.Println(data.(int64))
}

func doPanicWithParameter(input any) {
	defer util.CatchError()

	//? Test panic
	log.Println(input.(bool))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	doPanic()
	doPanicWithParameter("hello")

	for i := 0; i < 3; i++ {
		log.Println(i)
		time.Sleep(1 * time.Second)
	}
	log.Println("Program complete after multiple panic.")
}

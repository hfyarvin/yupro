package test_model

import(
	f "fmt"
	rt "runtime"
)

func say(s string)  {
	for i := 0; i < 5; i++ {
		rt.Gosched()
		f.Println(s)
	}
}

func GT()  {
	go say("world")
	say("hello")
}
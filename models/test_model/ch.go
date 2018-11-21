package test_model
// 有缓存channel与无缓存channel
import (
	f "fmt"
)

func sum(a []int, c chan int)  {
	total := 0
	for _, v := range a{
		f.Println("v: ", v)
		total += v
	}
	f.Println("f: ", total)
	c <- total
}

func Ch()  {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	f.Println(x, y, x+y)
}

// Buffered Channels, 先进先出
func BufCh()  {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	f.Println(<-c)
	f.Println(<-c)
}

// range和close
//通过语法v, ok := <-ch测试channel是否被关闭.如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭.
 func fibonacci(n int, c chan int)  {
	 x, y := 1, 1
	 for i := 0; i < n; i++ {
		 c <- x
		 x, y = y, x+y
	 }
	 close(c)
 }

 func ChClose()  {
	 c := make(chan int, 10)
	 go fibonacci(cap(c), c)

	 for i := range c {
		 f.Println(i)
	 }
 }

//  select: 监听channel上的数据流动
func fibonacci2(c, quit chan int)  {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			f.Println("quit")
			return
		default:
		}
	}
}

func Sel()  {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			f.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}
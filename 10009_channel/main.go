package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
channel 分位 buffered 和 unbufferd 2 种情况

几个问题
1. 为什么会有 deadlock 报错, 阻塞是如何实现的呢 ？
2.

channel 的使用模式

channel
buf
lock
sendx
recvx
sendq
recvq
qcount
elemsize
elemtype
dataqsiz
closed


接受
for  / for range / select
*/
func main() {
	a := 2
	a = 4
	a, b := 3, 2
	fmt.Print(a, " ",b)

	//WhenToCloseChannel()
	//anotherCaseToIllucaseSyncBetweenChannels()
	//syncBetweenChannel()
	//Channel123()

	//
	//for i:=0;i<3;i++ {
	//	fmt.Printf("i p %p, v:=%d\n", &i, i)
	//	testI(i)
	//}
}

func ForRangeAndSelect() {

}

func testI(i int) {
	fmt.Printf("i p %p, v:=%d\n", &i, i) // 不同的地址
}

// todo
// ..
func Channel123() {
	s := []string{"a", "b", "c", "d", "e", "f"}
	for _, v := range s {
		fmt.Printf("group 1 ,v address:%p type=%T \n", &v, v)
		go func(str string) {
			fmt.Printf("group 1 %p, %s\n", &v, v)
		}(v)
		//time.Sleep(time.Duration(time.Millisecond))
	}

	//for _, v := range s {
	//	tmpV:= v
	//	fmt.Printf("group 1 ,v address:%p\n", &v)
	//	go func(str string) {
	//		fmt.Printf("group 1 %p, %s\n", &v, v)
	//		fmt.Printf("group 1 tmp v %p, %s\n", &tmpV, tmpV)
	//	}(v)
	//	//time.Sleep(time.Duration(time.Millisecond))
	//}
	//
	//for i := 0; i < 3; i++ {
	//	go func(ii int) {
	//		fmt.Printf("output int %p, %d\n", &i, i)
	//	}(i)
	//}
	//
	//c := make(chan string, 10)
	//for _, x := range s {
	//	c <- x
	//}
	//close(c)
	//
	//for vv := range c {
	//	fmt.Printf("channel %p, %s\n", &vv, vv)
	//}

	//runtime.GOMAXPROCS(10)
	//runtime.Gosched()
	time.Sleep(time.Duration(time.Millisecond * 100))

}

// sync5
const NQ = 3

type MyLock struct {
	sync.Mutex
}

var (
	lock   MyLock
	status [NQ][2]bool
)

func firstJob(id int) {
	r := time.Duration(rand.Float64() * 100)
	time.Sleep(time.Duration(time.Millisecond * r))
	fmt.Printf("group %d have done job 1\n", id)
	status[id] = [2]bool{true, false}
}

func secondJob(id int) {
	for {
		lock.Lock()
		s := status[id]
		lock.Unlock()
		if s == [2]bool{true, false} {
			fmt.Printf("group %d have done job 2\n", id)
			status[id] = [2]bool{true, true}
			break //
		}
	}
}

func thirdJob(id int) {
	for {
		lock.Lock()
		s := status[id]
		lock.Unlock()
		if s == [2]bool{true, true} {
			fmt.Printf("group %d have done job 3\n", id)
			status[id] = [2]bool{true, true}
			break //
		}
	}
}

func syncBetweenChannel() {
	for i := 0; i < 3; i++ {
		go firstJob(i)
		go secondJob(i)
		go thirdJob(i)
	}

	time.Sleep(time.Duration(time.Second * 2))
}

// sync4
func anotherCaseToIllucaseSyncBetweenChannels() {
	fmt.Println("case to illustrate we need to sync between channels")
	c := make(chan string)
	defer close(c)
	go _sendValu(c, 1)
	go _sendValu(c, 2)
	v := <-c
	fmt.Printf("in main goroutine, receive v :%s\n", v)
	// if master exit before other goroutine, all channel will be destroyed, and data will lost
}

// sync3
func _sendValu(c chan string, id int) {
	fmt.Printf("executing goroutine from id:%d\n", id)
	time.Sleep(time.Second)
	c <- fmt.Sprintf("hello world from %d", id)
	fmt.Printf("finish executing goroutine from id:%d\n", id)
}

// sync2
func WhenToCloseChannel() {
	ch8 := make(chan int, 8)
	go func(c chan int) {
		//wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			//wg.Add(1)
			go func(c chan int, x int) {
				//defer wg.Done()
				time.Sleep(time.Duration(time.Millisecond * 10))
				c <- x // send to closed channel, so how to communicate between channel , wait group
			}(c, i)
		}
		//wg.Wait()
		close(ch8)
	}(ch8)

	time.Sleep(time.Duration(time.Second))
	for v := range ch8 {
		fmt.Println(v)
	}
}

// sync 1
func ForRangeWithChannel() {
	// Q7
	// for -- range will keep recv if there any more comming
	ch7 := make(chan int, 7)
	go func(c chan int) {
		for i := 0; i < 7; i++ {
			c <- i
		}
	}(ch7)
	for v := range ch7 {
		fmt.Println(v) // will deadklock because channel not close and keep receiving , and when channel is empty,
		// a dead lock error will occure
	}
}

func Q6() {
	//Q6
	ch6 := make(chan int, 5)
	ch6 <- 1
	ch6 <- 2
	// send to a full will block
	fmt.Println(<-ch6) //
	fmt.Println(<-ch6) //
	fmt.Println(<-ch6) // read a empty will block
	return
}

func Q5() {
	ch5 := make(chan int)
	fmt.Println(<-ch5) //  buffered
	go func(c chan int) {
		i := 0
		for {
			if i > 100 {
				break
			}
			c <- i
			i++
		}
	}(ch5)
	for v := range ch5 {
		fmt.Println("Q5", v)
	}
}

func Q4() {
	// Q4 buffered 的 channel 不会阻塞
	ch4 := make(chan int, 3)
	ch4 <- 1
	fmt.Printf("Q3 v:=%d", <-ch4)
	return
}

func Q3() {
	// Q3 unbuffered channel 会阻塞
	ch2 := make(chan int)
	ch2 <- 0
	fmt.Printf("Q3 v:=%d", <-ch2)
	return

}

// unbuffered channel, read from an empty buffered channel will block , send to a full also will block
func Q2() {
	// Q2
	ch1 := make(chan int, 3)
	go func(c chan int) {
		i := 0
		for {
			if i > 4 {
				break
			}
			c <- i
			time.Sleep(time.Duration(time.Millisecond * 10))
			i++
		}
	}(ch1)
	time.Sleep(time.Duration(time.Millisecond * 1000))
	close(ch1)

	//for i := 0; i < 10; i++ {
	//	v, ok := <-ch1
	//	fmt.Printf("%d: v:=%d, ok=%t\n", i, v, ok)
	//}
	j := 0
	for v := range ch1 {
		fmt.Printf("%d: v:=%d\n", j, v)
		j++
	}
}

// read from closed channel not raise panic,, will return 0, false
// if a channel if buffered and still have data not received , will read first then close
//
// write to a closed channel will raise panic
func Q1() {
	// Q1
	// unbuffer channel, 读取 closed channel 可以读取到内容, 向 closed channel 会 raised panic
	ch := make(chan int)
	go func(c chan int) {
		i := 0
		for {
			c <- i
			i++
		}
	}(ch)

	close(ch)

	for i := 0; i < 5; i++ {
		v, ok := <-ch
		// 从
		fmt.Printf("%d: v:=%d, ok=%t\n", i, v, ok)
	}
}

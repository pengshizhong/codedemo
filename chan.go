package main

import (
	"fmt"
)

func main()  {
	//chan1 := make(chan int, 1)
	//chan2 := chan1
	//fmt.Println(chan1 == chan2)
	////close(chan1)
	////chan1 <- 10
	//tmp, ok := <- chan1
	//fmt.Println(tmp, ok)
	//close(chan1)
	//tmp, ok = <- chan1
	//fmt.Println(tmp, ok)
	test()
}

//模拟ctx的替换closechan1
func test() {
	var closedchan = make(chan struct{})
	close(closedchan)
	//chan1 := make(chan int, 1)
	//chan2 := chan1
	//fmt.Println(chan1 == chan2)
	//close(chan1)
	//go print(chan1)
	//chan1 <- 10
	//tmp, ok := <- chan1
	//fmt.Println(tmp, ok)
	//close(chan1)
	//tmp, ok = <- chan1
	//fmt.Println(tmp, ok)
	//time.Sleep(1 * time.Second)
	//chan1 = closedchan
	//time.Sleep(1 * time.Second)
	fmt.Println("xiechen start")
	tmp, ok := <- closedchan
	fmt.Println(tmp, ok)
}

func print(chan1 chan int)  {
	for  {
		fmt.Println("xiechen start")
		tmp, ok := <- chan1
		fmt.Println(tmp, ok)
		fmt.Println("xiechen end")
	}

}
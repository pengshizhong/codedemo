package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main()  {
	fmt.Println("协程数量1：%d", runtime.NumGoroutine())
	//deadline()
	timeout()
	//time.Sleep(3*time.Second)
	fmt.Println("携程树林2：%d", runtime.NumGoroutine())

}

//cancel ctx的演示
func cancel() {
	//要知道ctx里面是怎么控制超时的
	//要知道http请求里面是怎么控制超时的
	//要熟练知道管道的作用
	rootCtx := context.TODO()
	pCtx, pcancel  := context.WithCancel(rootCtx)
	fmt.Println("----------第一个cancel初始化完成")
	ppCtx := context.WithValue(pCtx, "3", "2")
	fmt.Println("----------第一个value初始化完成")
	//_, pcancel  := context.WithCancel(rootCtx)
	defer pcancel()
	fmt.Println("valuectx地址：%v", &ppCtx)
	pppCtx, cancel := context.WithCancel(ppCtx)
	fmt.Println("----------第2个cancel初始化完成")
	cancel()
	fmt.Println(fmt.Sprintf("已经结束的管道输出值： %v", <- pppCtx.Done()))
	fmt.Println("----------完成第二个ctx取消")
	//go func(ctx context.Context) {
	//	// 发送HTTP请求
	//	done := make(chan bool, 1)
	//	for  {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("强制结束")
	//			return
	//		case <-done:
	//			fmt.Println("自然完成")
	//			return
	//		}
	//	}
	//}(ctx)
	//for  {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("call successfully!!!")
	//		return
	//	case <-time.After(time.Duration(time.Second*1)):
	//		cancel()
	//		fmt.Println("timeout!!!")
	//		return
	//	}
	//}
}

//deadline ctx的演示
func deadline() {
	//要知道ctx里面是怎么控制超时的
	//要知道http请求里面是怎么控制超时的
	//要熟练知道管道的作用
	rootCtx := context.TODO()
	pCtx, pcancel  := context.WithDeadline(rootCtx, time.Now())
	fmt.Println("----------第一个cancel初始化完成")
	ppCtx := context.WithValue(pCtx, "3", "2")
	fmt.Println("----------第一个value初始化完成")
	//_, pcancel  := context.WithCancel(rootCtx)
	defer pcancel()
	fmt.Println("valuectx地址：%v", &ppCtx)
	pppCtx, cancel := context.WithDeadline(ppCtx, time.Now())
	fmt.Println("----------第2个cancel初始化完成")
	cancel()
	fmt.Println(fmt.Sprintf("已经结束的管道输出值： %v", <- pppCtx.Done()))
	fmt.Println("----------完成第二个ctx取消")
	//go func(ctx context.Context) {
	//	// 发送HTTP请求
	//	done := make(chan bool, 1)
	//	for  {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("强制结束")
	//			return
	//		case <-done:
	//			fmt.Println("自然完成")
	//			return
	//		}
	//	}
	//}(ctx)
	//for  {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("call successfully!!!")
	//		return
	//	case <-time.After(time.Duration(time.Second*1)):
	//		cancel()
	//		fmt.Println("timeout!!!")
	//		return
	//	}
	//}
}

func timeout() {
	//要知道ctx里面是怎么控制超时的
	//要知道http请求里面是怎么控制超时的
	//要熟练知道管道的作用
	rootCtx := context.TODO()
	pCtx, pcancel  := context.WithTimeout(rootCtx, 1*time.Second)
	fmt.Println("----------第一个cancel初始化完成")
	ppCtx := context.WithValue(pCtx, "3", "2")
	fmt.Println("----------第一个value初始化完成")
	//_, pcancel  := context.WithCancel(rootCtx)
	defer pcancel()
	fmt.Println("valuectx地址：%v", &ppCtx)
	pppCtx, cancel := context.WithTimeout(ppCtx, 1*time.Second)
	fmt.Println("----------第2个cancel初始化完成")
	cancel()
	fmt.Println(fmt.Sprintf("已经结束的管道输出值： %v", <- pppCtx.Done()))
	fmt.Println("----------完成第二个ctx取消")
	//go func(ctx context.Context) {
	//	// 发送HTTP请求
	//	done := make(chan bool, 1)
	//	for  {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("强制结束")
	//			return
	//		case <-done:
	//			fmt.Println("自然完成")
	//			return
	//		}
	//	}
	//}(ctx)
	//for  {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("call successfully!!!")
	//		return
	//	case <-time.After(time.Duration(time.Second*1)):
	//		cancel()
	//		fmt.Println("timeout!!!")
	//		return
	//	}
	//}
}
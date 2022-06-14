package main

import "fmt"

type rsp struct {
	Ocode string
}

func main()  {
	//a := 2
	//b := &a
	//fmt.Println(b)
	c := &rsp{}
	d := &rsp{}
	fmt.Println("before", &c)
	fmt.Println("before", &d)
	fun1(c)
	fun2(d)
	fmt.Println(fmt.Sprintf("c: %v", c))
	fmt.Println(fmt.Sprintf("d: %v", d))
	fmt.Println(&c)
	fmt.Println(&d)
}

func fun1(r *rsp)  {
	r.Ocode = "fun1";
}

func fun2(r *rsp)  {
	fmt.Println("fun", &r)
	r = &rsp{
		"fun2",
	}
	fmt.Println("fun", &r)
	//r.Ocode = "fun2";
}
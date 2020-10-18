package main

import "fmt"

// func plus(x, y int) int {
// 	return x + y
// }
// func hello() {
// 	fmt.Println("Hello, World!")
// 	return
// }

// func div(a, b int) (int, int) {
// 	q := a / b
// 	r := a % b
// 	return q, r
// }

// func doSomething() (a int) {
// 	return
// }
// func doSomething() (x, y int) {
// 	y = 5
// 	x = 3
// 	return
// }

// func plus(x, y int) int {
// 	return x + y
// }

// func returnFunc() func() {
// 	return func() {
// 		fmt.Println("I'm a function")
// 	}
// }

// func callFunction(f func()) {
// 	f()
// }

// func later() func(string) string {
// 	var store string
// 	return func(next string) string {
// 		s := store
// 		store = next
// 		return s
// 	}
// }

// func integers() func() int {
// 	i := 0
// 	return func() int {
// 		i += 1
// 		return i
// 	}
// }

// const ONE = 1

// func one() (int, int) {
// 	const TWO = 2
// 	return ONE, TWO
// }

// func runDefer() {
// 	defer fmt.Println("defer")
// 	fmt.Println("done")
// }
// func testRecover(src interface{}) {
// 	defer func() {
// 		if x := recover(); x != nil {
// 			switch v := x.(type) {
// 			case int:
// 				fmt.Printf("panic: int=%v\n", v)
// 			case string:
// 				fmt.Printf("panic: string=%v\n", v)
// 			default:
// 				fmt.Println("panic: unknown")
// 			}
// 		}
// 	}()
// 	panic(src)
// 	return
// }
// func sub() {
// 	for {
// 		fmt.Println("sub loop")
// 	}
// }

var S = ""

func init() {
	S = S + "A"
}

func init() {
	S = S + "B"
}

func init() {
	S = S + "C"
}
func main() {
	fmt.Println(S)
	// go fmt.Println("Yeah!")
	// fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	// fmt.Printf("NumGorutine: %d\n", runtime.NumGoroutine())
	// fmt.Printf("Version: %s\n", runtime.Version())
	// go sub()
	// for {
	// 	fmt.Println("main loop")
	// }
	// testRecover(128)
	// testRecover("hogehoge")
	// testRecover([...]int{1, 2, 3})
	// defer fmt.Println("on defer")
	// panic("runtime error!")
	// fmt.Println("Hello, World!")
	// runDefer()
	// 	fmt.Println("A")
	// 	goto L
	// 	fmt.Println("B")
	// L:
	// 	fmt.Println("C")
	// var x interface{} = 3.14
	// switch v := x.(type) {
	// case bool:
	// 	fmt.Println("bool:", v)
	// case int:
	// 	fmt.Println(v * v)
	// case string:
	// 	fmt.Println(v)
	// default:
	// 	fmt.Printf("%#v\n", v)
	// }
	// var x interface{} = 3
	// switch x.(type) {
	// case bool:
	// 	fmt.Println("bool")
	// case int, uint:
	// 	fmt.Println("integer or unsigned integer")
	// case string:
	// 	fmt.Println("string")
	// default:
	// 	fmt.Println("don't know")
	// }
	// x, y := one()
	// fmt.Printf("x=%d, y=%d\n", x, y)
	// ints := integers()
	// fmt.Println(ints())
	// fmt.Println(ints())
	// fmt.Println(ints())
	// otherInts := integers()
	// fmt.Println(otherInts())
	// f := later()
	// fmt.Println(f("Golang"))
	// fmt.Println(f("is"))
	// fmt.Println(f("awesome"))
	// callFunction(func() {
	// 	fmt.Println("I'm a function")
	// })
	// f := returnFunc()
	// f()
	// returnFunc()()
	// var plusAlias = plus

	// fmt.Printf("%v", plusAlias(10, 5))
	// f := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Printf("%d\n", f(2, 3))
	// fmt.Printf("%T\n", func(x, y int) int { return x + y })
	// fmt.Printf("%#v\n", func(x, y int) int { return x + y })
	// fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))
	// fmt.Printf("%v", plus(3, 4))
	// hello()
	// q, r := div(19, 7)
	// fmt.Printf("商=%d 剰余=%d\n", q, r)
	// fmt.Printf("%d, %d", doSomething())
}

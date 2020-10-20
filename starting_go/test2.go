package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main() {
	ch := make(chan int, 20)

	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i += 1
	}
	close(ch)
	time.Sleep(3 * time.Second)
	// ch := make(chan int, 3)

	// ch <- 1
	// ch <- 2
	// ch <- 3

	// close(ch)

	// var (
	// 	i  int
	// 	ok bool
	// )

	// i, ok = <-ch
	// i, ok = <-ch
	// i, ok = <-ch
	// i, ok = <-ch
	// fmt.Println(i, ok)
	// ch := make(chan rune, 3)

	// ch <- 'A'
	// ch <- 'B'
	// ch <- 'C'
	// ch <- 'D'
	// fmt.Println(ch)
	// ch := make(chan int)
	// go receiver(ch)

	// i := 0
	// for i < 10 {
	// 	ch <- i
	// 	i += 1
	// }
	// m := map[int]string{
	// 	1: "A",
	// 	2: "B",
	// 	3: "C",
	// }
	// fmt.Println(len(m))
	// m[4] = "D"
	// m[5] = "E"
	// fmt.Println(len(m))
	// delete(m, 4)
	// fmt.Println(m)
	// m := map[int]string{
	// 	1: "Apple",
	// 	2: "Banana",
	// 	3: "Cherry",
	// }

	// for k, v := range m {
	// 	fmt.Printf("%d -> %s\n", k, v)
	// }
	// m := map[int]string{1: "A", 2: "B", 3: "C"}
	// // s := m[1]
	// // fmt.Println(s)
	// s, ok := m[2]
	// fmt.Println(s, ok)

	// m := map[int]map[float64]string{
	// 	1: {3.14: "PI"},
	// }

	// m := map[int][]int{
	// 	1: {1},
	// 	2: {1, 2},
	// 	3: {1, 2, 3},
	// }
	// fmt.Println(m)
	// m := map[int]string{
	// 	1: "Taro",
	// 	2: "Hanako",
	// 	3: "Jiro",
	// }
	// fmt.Println(m)
	// m := make(map[string]string)

	// m["Yamada"] = "Taro"
	// m["Sato"] = "Hanako"
	// m["Yamada"] = "Jiro"
	// fmt.Println(m)
	// m := make(map[int]string)
	// m[1] = "US"
	// m[81] = "Japan"
	// m[86] = "China"
	// fmt.Println(m)
}

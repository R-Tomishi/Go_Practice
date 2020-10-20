package main

import "fmt"

// func sum(s ...int) int {
// 	n := 0
// 	for _, v := range s {
// 		n += v
// 	}
// 	return n
// }

// func pow(a []int) {
// 	for i, v := range a {
// 		a[i] = v * v
// 	}
// 	return
// }
func main() {
	a := [3]int{1, 2, 3}
	s := a[:]
	fmt.Println(len(s))
	fmt.Println(cap(s))
	a[0] = 8
	fmt.Println(a)
	fmt.Println(s)

	s = append(s, 4)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	a[1] = 11
	fmt.Println(a)
	fmt.Println(s)

	// a := []int{1, 2, 3}
	// pow(a)
	// fmt.Println(a)
	// s := []int{1, 2, 3, 4}
	// fmt.Println(sum(1, 2, 3))
	// fmt.Println(sum(1, 2, 3, 4, 5))
	// fmt.Println(sum())
	// fmt.Println(sum(s...))
	// // s := []string{"Apple", "Banana", "Cherry"}
	// for i, v := range s {
	// 	fmt.Printf("[%d] -> %s\n", i, v)
	// 	s = append(s, "Melon")
	// }
	// fmt.Println(s)
	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// s1 := a[2:4]
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

	// s2 := a[2:4:4]
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))

	// s3 := a[2:4:6]
	// fmt.Println(len(s3))
	// fmt.Println(cap(s3))
	// b := make([]byte, 9)
	// n := copy(b, "あいうえお")
	// fmt.Println(n, b)
	// s1 := []int{1, 2, 3, 4, 5}
	// s2 := []int{10, 11, 9, 8, 7, 6}
	// copy(s1, s2)
	// // fmt.Println(n)
	// fmt.Println(s1)
	// fmt.Println(s2)
	// s := make([]int, 1024, 1024)
	// fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))

	// s = append(s, 0)
	// fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	// s := make([]int, 0, 0)
	// fmt.Printf("(A) len=%d, cap=%d\n", len(s), cap(s))

	// s = append(s, 1)
	// fmt.Printf("(B) len=%d, cap=%d\n", len(s), cap(s))

	// s = append(s, []int{2, 3, 4}...)
	// fmt.Printf("(C) len=%d, cap=%d\n", len(s), cap(s))

	// s = append(s, 5)
	// fmt.Printf("(D) len=%d, cap=%d\n", len(s), cap(s))

	// s = append(s, 6, 7, 8, 9)
	// fmt.Printf("(E) len=%d, cap=%d\n", len(s), cap(s))

	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// a := s[:2]
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))
	// s = append(s, 4)
	// fmt.Println(s)
	// s = append(s, 1, 2, 3)
	// fmt.Println(s)
	// s1 := append(s, a...)
	// fmt.Println(s1)

	// var b []byte
	// var c []string
	// b = append(b, "あいうえお"...)
	// b = append(b, "かきくけこ"...)
	// b = append(b, "さしすせそ"...)
	// c = append(c, "あいうえお")
	// c = append(c, "かきくけこ")
	// c = append(c, "さしすせそ")

	// fmt.Println(b)
	// fmt.Println(c)
	// s1 := make([]int, 5)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))
	// s2 := make([]int, 5, 10)
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))
	// a := make([]float64, 3)
	// fmt.Println(a)
	// a[0] = 3.14
	// fmt.Println(a)
	// a[1] = 6.28
	// fmt.Println(a)
	// fmt.Println(a[0])

	// s := make([]int, 8)
	// fmt.Println(len(s))
	// fmt.Println(a[4])
	// s := make([]int, 10)
	// var a [10]int
	// fmt.Println(s)
	// fmt.Println(a)
}

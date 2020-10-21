package main

import (
	"fmt"
	"math"
)

// func inc(p *int) {
// 	*p += 1
// }

// func pow(p *[3]int) {
// 	i := 0
// 	for i < 3 {
// 		(*p)[i] = (*p)[i] * (*p)[i]
// 		i += 1
// 	}
// }

// type MyInt int

// var n1 MyInt = 5

// type (
// 	IntPair	[2]int
// 	Strings []string
// 	AreaMap map[string][2]float64
// 	IntsChannel chan []int
// )

// type Callback func(i int) int

// func Sum(ints []int, callback Callback) int {
// 	var sum int
// 	for _, i := range ints {
// 		sum += i
// 	}
// 	return callback(sum)
// }

// type Point struct {
// 	X int
// 	Y int
// }

// var pt Point

// type Feed struct {
// 	Name   string
// 	Amount uint
// }

// type Animal struct {
// 	Name string
// 	Feed
// }

// type T0 struct {
// 	Name1 string
// }

// type T1 struct {
// 	T0
// 	Name2 string
// }

// type T2 struct {
// 	T1
// 	Name3 string
// }

// type Base struct {
// 	Id    int
// 	Owner string
// }

// type A struct {
// 	Base
// 	Name string
// 	Area string
// }

// type B struct {
// 	Base
// 	Title  string
// 	Bodies []string
// }

type Point struct {
	X, Y int
}

// func (p *Point) Render() {
// 	fmt.Printf("<%d, %d>\n", p.X, p.Y)
// }

// func swap(p *Point) {
// 	x, y := p.Y, p.X
// 	p.X = x
// 	p.Y = y
// }

// type Person struct {
// 	Id   int
// 	Name string
// 	Area string
// }

func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	p := &Point{X: 0, Y: 0}
	fmt.Println(p.Distance(&Point{X: 1, Y: 1}))

	// p := new(Point)
	// p.X = 1
	// p.Y = 2
	// p := new(Person)
	// fmt.Println(p.Id, p.Name, p.Area)
	// p := Point{X: 1, Y: 2}
	// swap(&p)
	// fmt.Println(p.X, p.Y)
	// a := A{
	// 	Base: Base{17, "Taro"},
	// 	Name: "Taro",
	// 	Area: "Tokyo",
	// }
	// b := B{
	// 	Base:   Base{81, "Hanako"},
	// 	Title:  "no title",
	// 	Bodies: []string{"A", "B"},
	// }
	// fmt.Println(a.Id, a.Owner, b.Id, b.Owner)
	// t := T2{T1: T1{T0: T0{Name1: "X"}, Name2: "Y"}, Name3: "Z"}
	// fmt.Println(t.Name1)
	// fmt.Println(t.Name2)
	// fmt.Println(t.Name3)
	// a := Animal{
	// 	Name: "Monkey",
	// 	Feed: Feed{
	// 		Name:   "Banana",
	// 		Amount: 10,
	// 	},
	// }
	// fmt.Println(a.Name)
	// fmt.Println(a.Feed.Name)
	// fmt.Println(a.Amount)
	// a.Feed.Amount = 15
	// fmt.Println(a.Amount)
	// fmt.Println(pt.X)
	// fmt.Println(pt.Y)
	// pt.X = 10
	// pt.Y = 8
	// fmt.Println(pt.X)
	// fmt.Println(pt.Y)
	// pt := Point{1, 2}
	// fmt.Println(pt.X)
	// fmt.Println(pt.Y)
	// pt = Point{Y: 28}
	// fmt.Println(pt.X)
	// fmt.Println(pt.Y)

	// n := Sum(
	// 	[]int{1, 2, 3, 4, 5},
	// 	func(i int) int {
	// 		return i * 2
	// 	},
	// )
	// fmt.Println(n)
	// pair := IntPair{1, 2}
	// strs := Strings{"Apple", "Banana", "Cherry"}
	// amap := AreaMap{"Tokyo": {35.689488, 139.601706}}
	// ich := make(IntsChannel)
	// n2 := MyInt(7)
	// fmt.Println(n1)
	// fmt.Println(n2)
	// s := "ABC"
	// fmt.Println(s[1])
	// p := &[3]int{1, 2, 3}
	// pow(p)
	// fmt.Println(p[2])
	// i := 1
	// inc(&i)
	// fmt.Println(i)
	// inc(&i)
	// fmt.Println(i)
	// inc(&i)
	// fmt.Println(i)
	// var i int
	// p := &i
	// i = 5
	// fmt.Println(*p)
	// *p = 10
	// fmt.Println(i)
}

package main

import "fmt"

// type IntPair [2]int

// func (ip IntPair) First() int {
// 	return ip[0]
// }

// func (ip IntPair) Last() int {
// 	return ip[1]
// }

// type Strings []string

// func (s Strings) Join(d string) string {
// 	sum := ""
// 	for _, v := range s {
// 		if sum != "" {
// 			sum += d
// 		}
// 		sum += v
// 	}
// 	return sum
// }

// type User struct {
// 	Id   int
// 	Name string
// }

// func NewUser(id int, name string) *User {
// 	u := new(User)
// 	u.Id = id
// 	u.Name = name
// 	return u
// }

// type Point struct {
// 	X, Y int
// }

// func (p *Point) ToString() string {
// 	return fmt.Sprintf("[%d, %d]", p.X, p.Y)
// }

// func (p *Point) Set(x, y int) {
// 	p.X = x
// 	p.Y = y
// }

// type Point struct {
// 	X, Y int
// }

// type Points []*Point

// func (ps Points) ToString() string {
// 	str := ""
// 	for _, p := range ps {
// 		if str != "" {
// 			str += ","
// 		}
// 		if p == nil {
// 			str += "<nil>"
// 		} else {
// 			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
// 		}
// 	}
// 	return str
// }

// type User struct {
// 	Id   int
// 	Name string
// }

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

func main() {
	vs := []Stringify{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "XXX-0123", Model: "PX512"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	Println(&Person{Name: "Hanako", Age: 23})
	Println(&Car{Number: "XYZ-9999", Model: "RT-3"})
	// m1 := map[User]string{
	// 	{Id: 1, Name: "Taro"}: "Tokyo",
	// 	{Id: 2, Name: "Jiro"}: "Osaka",
	// }
	// m2 := map[int]User{
	// 	1: {Id: 1, Name: "Taro"},
	// 	2: {Id: 2, Name: "Jiro"},
	// }
	// ps := Points{}
	// fmt.Println(ps)
	// ps = append(ps, &Point{X: 1, Y: 2})
	// fmt.Println(ps)
	// ps = append(ps, nil)
	// ps = append(ps, &Point{X: 3, Y: 4})
	// fmt.Println(ps)
	// fmt.Println(ps.ToString())
	// ps := make([]Point, 5)
	// for _, p := range ps {
	// 	fmt.Println(p.X, p.Y)
	// }
	// t := &foo.T{}
	// fmt.Println(t.Method1())
	// fmt.Println(t.Field1)

	// fmt.Println(t.method2())
	// fmt.Println(t.field2)
	// p1 := Point{}
	// p1.Set(1, 2)
	// fmt.Println(p1.X)
	// fmt.Println(p1.Y)

	// p2 := &Point{}
	// p2.Set(1, 2)
	// fmt.Println(p2.X)
	// fmt.Println(p2.Y)
	// f := (*Point).ToString
	// fmt.Println(f(&Point{X: 7, Y: 11}))
	// fmt.Println(NewUser(1, "Taro"))
	// ip := IntPair{1, 2}
	// fmt.Println(ip.First())
	// fmt.Println(ip.Last())

	// fmt.Println(Strings{"A", "B", "C"}.Join(","))
}

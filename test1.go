package main

import (
	"fmt"
	"strings"
	"math"
)


func section1() {

	// point
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)

	*p = 21

	fmt.Println(i)
}

func section2() {

	// struct

	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	fmt.Println(v.X)
}

func section3() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	var x *Vertex = &v
	fmt.Println(*x)

	fmt.Println(p)
	fmt.Println(v)
}

func section4() {

	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{X: 1}  // Y:0 被省略
		v3 = Vertex{}      // X:0 和 Y:0
		p  = &Vertex{1, 2} // 类型为 *Vertex
	)

	fmt.Println(v1, v2, v3, p)
}

func section5() {

	p := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("p===", p)

	for i:=0; i<len(p); i++ {
		fmt.Printf("p[%d]== %d \n", i, p[i])
	}
}

func section6() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p== ", p)
	fmt.Println("p[1:4]== ", p[1:4])
	fmt.Println("p[:3]== ", p[:3])
	fmt.Println("p[4:]", p[4:])

}

func section7() {
	a := make([]int, 5)

	fmt.Println(a, len(a), cap(a))

	b := make([]int, 0, 5)
	fmt.Println(b, b[:5])

	/* 这里有疑问，这个内存分配的问题是怎么分配的？
		对于下面的切片的赋值，其实只是将原来的切片的中的哪个位置地址给与新变量的起始位置，
		能力时以新变量的当前地址位置，截止到原来切片的末端，
		参考地址：http://stackoverflow.com/questions/12768744/re-slicing-slices-in-golang

		v := make([]int, 5, 5)
		printSring(v)

		a := v[:2]
		printSring(a)

		b := v[1:5]
		printSring(b)

		c := b[:2]
		printSring(c)
		c[0] = 1
		printSring(c)
		printSring(v)
	*/

	c := b[:2]
	fmt.Println(c, len(c), cap(c))


	d := b[2: 5]
	fmt.Println(d, len(d), cap(d))

	//e := d[:3]
	//e = append(e, 0, 1, 6, 7)
	//fmt.Println(e, len(e), cap(e))
	//fmt.Println(b[1:])
}

func section8() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func section9()  {
	var a []int

	fmt.Println(a, len(a), cap(a))

	a = append(a, 0)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 2, 3, 4, 5)
	fmt.Println(a, len(a), cap(a))

}


func printSring(x []int) {
	fmt.Printf("slice: %v, len=%d, cap=%d\n", x, len(x), cap(x))

	//fmt.Println(&x)
}

func section10() {
	var pow = []int {1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2^%d = %d \n", i, v)
	}

}

func section11() {

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	fmt.Println(pow)
	for _, value := range pow {
		fmt.Printf("%d \n", value)
	}
}

func section12() {
	type Vertex struct {
		Lat, Long float64
	}
    //
	var m map[string]Vertex

	m = make(map[string]Vertex)

	m["bell labs"] = Vertex{40.68433, -743.39967}

	fmt.Println(m["bell labs"])
}

func section13() {

	var value = map[string]int {
		"id": 1,
		"age": 2,
	}

	fmt.Println(value["id"], value["age"])
}

func section14() {

	m := make(map[string]int)

	m["age"] = 42
	fmt.Println(m["age"])

	delete(m, "age")
	fmt.Println(m["age"])

	value, exist := m["age"]

	fmt.Println(value, exist)

}


func section15() {
	fmt.Printf("Fields are: %q", strings.Fields("foo bar baz"))
}


func section16(s string) map[string]int {
	values := strings.Fields(s)

	response := make(map[string]int)
	for _, v := range values {
		if len(v) == 1 {
			response[v] = strings.Count(s, v + " ")
		} else {
			response[v] = strings.Count(s, v)
		}

	}
	return response
}

func section17(x, y float64) float64 {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}
	return hypot(x, y)
}


func addr() func(int) int {
	sum := 0

	return func(x int) int {
		fmt.Printf("sum=%d   ", sum)
		sum += x
		return sum
	}
}

func section18()  {
	pos, _ := addr(), addr()
	for i:=0; i<10; i++ {
		fmt.Println(
			pos(i),
			//neg(-2 * i),
		)
	}
}


func fibonacci() func() int {

	sum := -1
	var _fn, __fn int
	return func() int {
		sum += 1
		switch sum {
		case 0:
			return 0
		case 1:
			_fn = sum
			return 1
		default:
			sum = _fn + __fn
			__fn = _fn
			_fn = sum
			return sum
		}

	}
}

func section19() {
	f := fibonacci()
	for i:=0; i<10; i++ {
		fmt.Println(f())
	}
}

func main() {

	//fmt.Println(section16("I ate a donut. Then I ate another donut."))

	//fmt.Println(section17(3, 4))

	section19()


}

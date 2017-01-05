package main

import "fmt"
import (
	"reflect"
	"math"
	"time"
)

func add(x int, y int) int {

	return x + y
}

func del(a, b int) float32 {

	return float32(a - b)
}

func swap(x, y string) (string, string) {

	return "you" + y, "me" + x
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {

		return v
	}
	return lim
}

func Sqrt(x float64) float64 {

	var z, _z float64 = 1.0, 1.0
	for {
		z = z - (z * z - x) / (2 * z)

		v := _z - z
		if v < 0 {
			v = -v
		}
		if v < 0.00000000001 {
			break
		}
		_z = z

	}

	return z

}


func main() {
	fmt.Printf("hello, world.\n")

	fmt.Println(add(2, 4))

	fmt.Println(del(3, 2))

	name1, name2 := swap("hello", "world")
	fmt.Println(name1, name2)

	var num1, num2 string
	num1 = "joker"
	num2 = "laddy"
	fmt.Println(num1, num2)

	var u1, u2 = 1, 2

	fmt.Println(u1, u2)

	f := 3.1415926
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	for i:=1; i<10; i++{
		fmt.Println(i)
	}


	fmt.Println(sqrt(2), sqrt(-4))


	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	if true {
		fmt.Println(456)
	} else {
		fmt.Println("123")
	}

	fmt.Println(Sqrt(125348))

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}



	defer fmt.Println("world")
	fmt.Println("hello")


	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}


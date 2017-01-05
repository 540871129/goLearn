package main


import (
    "fmt"
    "math"
    "os"
    "time"
    "strings"
    "io"
    "net/http"
    "log"
)

type Vertex struct {
    X, Y float64
}


func(v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {

    v.X = v.X * f
    v.Y = v.Y * f
}

func x_section() {


    v := &Vertex{3, 4}
    fmt.Println(v.Abs())

}

func x_section1()  {

    v := &Vertex{3, 4}
    u := Vertex{3, 4}
    x := Vertex{3, 4}
    v.Scale(5)
    u.Scale(5)


    fmt.Println(v, v.Abs())
    fmt.Println(u, u.Abs())
    fmt.Println(x, x.Abs())

}


type Abser interface {
    Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {

    if f < 0 {
        f = -f
    }
    return float64(f)
}


func x_section3() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f
    a = &v
    //a = v   // error
    fmt.Println(a.Abs())
}


type Reader interface {
    Read(b []byte) (n int, err error)
}

type Writer interface {
    Write(b []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

func x_section4()  {
    var w ReadWriter
    w = os.Stdout
    fmt.Fprintf(w, "hello, writer\n")
}


type Person struct {
    Name string
    Age int
}

func (p Person) String() string {

    return fmt.Sprintf("%v (%v) years", p.Name, p.Age)
}


func x_section5() {

    a := Person{"Arthur Dent", 42}
    z := Person{"Joker", 23}
    fmt.Println(a, z)

}


type IPAddress [4]byte

func (ip IPAddress) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])

}


func x_section6() {

    address := map[string]IPAddress {
        "lookback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for n, a := range address {
        fmt.Printf("%v: %v\n", n, a)
    }
}

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error  {

    return &MyError{time.Now(), "it don't work",}
}

func x_section7() {
    if err := run(); err != nil {
        fmt.Println(err)
    }

    var e error
    e = &MyError{time.Now(), "it donno"}
    fmt.Println(e)
}


type ErrNegativeSqrt float64


func (e ErrNegativeSqrt) Error() string {

    fmt.Println(e)
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}


func XSqrt(x float64) (float64, error) {

    if x < 0{
        return 0, ErrNegativeSqrt(x)
    }

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

    return z, nil
}

func x_section8() {

    if value, err := XSqrt(-2); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(value)
    }


}


func x_section9() {
    //r := strings.NewReader("Hello, Reader!")
    r := strings.NewReader("Lbh penpxrq gur pbqr!")

    b := make([]byte, 32768)

    for {
        n, err := r.Read(b)

        fmt.Printf("n=%v, err=%v b=%v \n", n, err, b)
        fmt.Printf("b[:n]=%q \n", b[:n])
        if err == io.EOF {
            break
        }
    }
}


type rot13Reader struct {
    r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error)  {

    /*
        这里的b长度有32768，每次通过rot.r.Read(b)读取后，将r的数据读取到b切片中，相应的返回了读取的长度和错误，然后
        通过循环，来对b中的字符作一个变换，switch fallthrough 是直接越过后续一个判断条件然后操作，默认switch每个case的
        判断都是执行后直接break，不像java c会要显式break。io.copy会不断的调用rot13的Read方法，直到方法返回一个EOF错误
        才会停止copy
    */


    n, err := rot.r.Read(b)
    length := len(b)

    fmt.Println(length, n, err)

    for i:=0; i< length; i++ {
        switch  {

        case b[i] >= 'a' && b[i] < 'n':
            fallthrough
        case b[i] >= 'A' && b[i] < 'N' :
            b[i] += 13
        case b[i] >= 'n' && b[i] <= 'z':
            fallthrough
        case b[i] >= 'N' && b[i] <= 'Z':
            b[i] -= 13
        }

    }

    return n, err


}
// 这个练习不清楚需要怎么作
func x_section10() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}


type Hello struct {}


func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte{'h', 'e', 'l', 'l', 'o'})

    fmt.Println(r.Method, "some body request!")
}

func x_section11() {

    var h Hello

    err := http.ListenAndServe("localhost:4000", h)
    if err != nil {
        log.Fatal(err)
    }

}

func main()  {

    //x_section9()
    //x_section11()

    x_section10()



}

